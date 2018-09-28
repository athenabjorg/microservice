package mqConnection

import (
	"log"

	"github.com/streadway/amqp"
)

type Connection struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func OpenConnection() *Connection {
	c := &Connection{
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

func CloseConnection(c *Connection) {
	c.conn.Close()
	c.channel.Close()
}

func GetMessages(c *Connection) <-chan amqp.Delivery {
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

func SendMessage(c *Connection, msg []byte) {
	err := c.channel.Publish(
		"",      // exchange
		"hello", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", msg)
}
