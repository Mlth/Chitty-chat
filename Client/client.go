package main

import (
	"context"
	"fmt"
	"log"

	chat "github.com/Mlth/Chitty-chat/proto"
	"google.golang.org/grpc"
)

var name string = ""
var clock int32 = 0

func main() {
	// Create a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	fmt.Print("Write your name: ")
	fmt.Scan(&name)

	//  Create new Client from generated gRPC code from proto
	c := chat.NewChatClient(conn)

	go JoinServer(c)
	SendMessage(c)
}

func JoinServer(c chat.ChatClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	message := chat.WrittenMessage{Name: name}

	response, err := c.JoinServer(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	for {
		var responseMessage, _ = response.Recv()

		if responseMessage.Name == "" {
			log.Printf(responseMessage.Message)
		} else {
			log.Printf(responseMessage.Name + ": " + responseMessage.Message)
		}
	}
}

func SendMessage(c chat.ChatClient) {
	for {
		var inputMessage string
		fmt.Scan(&inputMessage)
		message := chat.WrittenMessage{Name: name, Message: inputMessage, TimeStamp: 0}

		_, err := c.SendMessage(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when sending message: %s", err)
		}
	}
}
