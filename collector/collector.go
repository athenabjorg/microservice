package main

import (
	"fmt"
	"log"
	"time"

	"github.com/athenabjorg/microservice/mqConnection"
	"github.com/athenabjorg/microservice/proto"

	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func main() {

	c := mqConnection.OpenConnection()
	defer mqConnection.CloseConnection(c)

	msgs := mqConnection.GetMessages(c)

	forever := make(chan bool)
	go processMessages(msgs)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

func processMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		m := &messagepb.Message{}

		err := proto.Unmarshal(d.Body, m)

		if err != nil {
			d.Nack(false, true) // (multiple, requeue)
			log.Fatalln("Failed to parse Message:", err)
		}

		log.Printf("Received a message: %s", m)

		// 	//if kill message
		// 	//processKillMessage()
		// 	//else
		// 	//saveMessageToRedis()
		time.Sleep(3 * time.Second)
		d.Ack(false) // (multiple)
	}
}

func processKillMessage() {

	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	// if err != nil {
	// 	log.Fatalf("Could not connect: %v", err)
	// }

	// defer conn.Close()

	// c := calculatorpb.NewCalculatorServiceClient(conn)

	// forwardKillMessageToAggregator(c)
	forwardKillMessageToAggregator()
}

// func forwardKillMessageToAggregator(c calculatorpb.CalculatorServiceClient) {
func forwardKillMessageToAggregator() {
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
}

func saveMessageToRedis() {
	fmt.Println("Saving message to redis")
}
