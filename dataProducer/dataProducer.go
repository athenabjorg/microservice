package main

import (
	"log"

	"github.com/athenabjorg/microservice/mqConnection"
	"github.com/athenabjorg/microservice/proto"

	"github.com/golang/protobuf/proto"
)

func main() {

	c := mqConnection.OpenConnection()
	defer mqConnection.CloseConnection(c)

	msgs := []*messagepb.Message{
		&messagepb.Message{
			Receiver:  "a",
			Sender:    "b",
			Value:     1,
			Operation: "sum",
		},
		&messagepb.Message{
			Receiver:  "a",
			Sender:    "b",
			Value:     2,
			Operation: "done",
		},
	}

	for _, msg := range msgs {
		protoMsg, err := proto.Marshal(msg)
		if err != nil {
			log.Fatalln("Failed to encode message:", err)
		}
		mqConnection.SendMessage(c, protoMsg)
	}
}
