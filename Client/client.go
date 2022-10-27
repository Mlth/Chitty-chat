package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	chat "github.com/Mlth/Chitty-chat/proto"
	"google.golang.org/grpc"
)

var name string = ""
var clock int32 = 0
var reader = bufio.NewReader(os.Stdin)

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
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//  Create new Client from generated gRPC code from proto
	c := chat.NewChatClient(conn)

	go JoinServer(c)
	SendMessage(c)
}

// This functions joins the server and keeps checking if it has recieved a message in its stream
func JoinServer(c chat.ChatClient) {
	//Increments clock because the client sends a message to the server that they want to join
	clock += 1
	message := chat.WrittenMessage{Name: name, TimeStamp: clock}

	response, err := c.JoinServer(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when trying to join server: %s", err)
	}

	for {
		var responseMessage, _ = response.Recv()
		//Increments clock again when the client recieves a message through its stream
		syncClock(responseMessage.TimeStamp)
		clock += 1

		if responseMessage.Name == "" {
			log.Printf(responseMessage.Message + ", timestamp: " + strconv.FormatInt(int64(clock), 10))
		} else {
			log.Printf(responseMessage.Name + ": " + responseMessage.Message + ", timestamp: " + strconv.FormatInt(int64(clock), 10))
		}
	}
}

// This function waits for the client to type a message in the console, and sends that mesage to the server
func SendMessage(c chat.ChatClient) {
	for {
		inputMessage, _ := reader.ReadString('\n')
		inputMessage = strings.TrimSpace(inputMessage)
		if len(inputMessage) > 128 {
			log.Println("Message should be below 128 characters, try again!")
			continue
		}
		//The clock is incremented because the user sends a message.
		clock += 1
		message := chat.WrittenMessage{Name: name, Message: inputMessage, TimeStamp: clock}
		_, err := c.SendMessage(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when sending message: %s", err)
		}
	}
}

func syncClock(timestamp int32) {
	if clock < timestamp {
		clock = timestamp
	}
}
