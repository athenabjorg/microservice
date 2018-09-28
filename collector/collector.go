package main

import (
	"fmt"
	"log"

	"github.com/athenabjorg/microservice/mqConnection"
	"github.com/athenabjorg/microservice/proto"
	"github.com/athenabjorg/microservice/redisConnection"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"
)

func main() {

	mqConn := mqconnection.OpenConnection()
	defer mqconnection.CloseConnection(mqConn)
	msgs := mqconnection.GetMessages(mqConn)

	redis := redisconnection.OpenConnection()
	defer redisconnection.CloseConnection(redis)

	forever := make(chan bool)
	for d := range msgs {
		go processMessage(d, redis)
	}
	fmt.Println(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever
}

func processMessage(d amqp.Delivery, redis redis.Conn) {

	msg := &messagepb.Message{}

	err := proto.Unmarshal(d.Body, msg)

	if err != nil {
		d.Nack(false, true) // (multiple, requeue)
		log.Fatalln("Failed to parse Message:", err)
	}

	fmt.Println("Received a message: ", msg)

	if msg.GetOperation() == "done" {
		// TODO return error and do nack if needed
		processDoneMessage()
	} else {
		// TODO return error and do nack if needed
		saveMessageToRedis(redis)
	}

	d.Ack(false) // (multiple)

}

func processDoneMessage() {

	// conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	// if err != nil {
	// 	log.Fatalf("Could not connect: %v", err)
	// }

	// defer conn.Close()

	// c := calculatorpb.NewCalculatorServiceClient(conn)

	// forwardDoneMessageToAggregator(c)
	forwardDoneMessageToAggregator()
}

// func forwardDoneMessageToAggregator(c calculatorpb.CalculatorServiceClient) {
func forwardDoneMessageToAggregator() {
	fmt.Println("Forwarding a done message")

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

func saveMessageToRedis(redis redis.Conn) {
	//set
	fmt.Println("Sending message to redis")
	redis.Do("SET", "message1", "Hello World")
}
