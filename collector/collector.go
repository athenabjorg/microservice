package main

import (
	"fmt"
	"log"
	"time"

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

func main() {

	c := &Consumer{
		conn:    nil,
		channel: nil,
	}
	var err error

	c.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer c.conn.Close()

	c.channel, err = c.conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer c.channel.Close()

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

	forever := make(chan bool)
	go processMessages(msgs)``
	<-forever
}

func processMessages(msgs <-chan amqp.Delivery) {

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// 	//if kill message
			// 	//processKillMessage()
			// 	//else
			// 	//saveMessageToRedis()
			time.Sleep(5 * time.Second)
			d.Ack(true)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// for {

	// }
}

func processKillMessage() error {

	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	// if err != nil {
	// 	log.Fatalf("Could not connect: %v", err)
	// }

	// defer conn.Close()

	// c := calculatorpb.NewCalculatorServiceClient(conn)

	// forwardKillMessageToAggregator(c)
	return forwardKillMessageToAggregator()
}

// func forwardKillMessageToAggregator(c calculatorpb.CalculatorServiceClient) bool {
func forwardKillMessageToAggregator() error {
	fmt.Println("Forwarding a kill message")

	// req := &calculatorpb.SumRequest{
	// 	Values: &calculatorpb.Values{
	// 		A: 5,
	// 		B: 7,
	// 	},
	// }

	// res, err := c.Sum(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("Error while calling Calculator RPC: %v", err)
	// 	return err
	// }

	// log.Printf("Response from Sum: %v", res.Result)
	return nil
}

func saveMessageToRedis() {
	fmt.Println("Saving message to redis")
}
