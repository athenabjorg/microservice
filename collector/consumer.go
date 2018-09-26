package main

import (
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CreateConsumer() *Consumer {
	c := &Consumer{
		conn:    nil,
		channel: nil,
	}
	var err error

	c.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	c.channel, err = c.conn.Channel()
	failOnError(err, "Failed to open a channel")

	return c
}

func CloseConsumer(c *Consumer) {
	c.conn.Close()
	c.channel.Close()
}

func GetConsumerMessages(c *Consumer) <-chan amqp.Delivery {
	msgs, err := c.channel.Consume(
		"hello", // queue
		"",      // consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer")

	return msgs
}
