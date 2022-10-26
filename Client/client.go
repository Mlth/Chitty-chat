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

func JoinServer(c chat.ChatClient) {
	// Between the curly brackets are nothing, because the .proto file expects no input.
	clock += 1
	message := chat.WrittenMessage{Name: name, TimeStamp: clock}

	response, err := c.JoinServer(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling GetTime: %s", err)
	}

	for {
		var responseMessage, _ = response.Recv()
		syncClock(responseMessage.TimeStamp)
		clock += 1

		if responseMessage.Name == "" {
			log.Printf(responseMessage.Message + ", timestamp: " + strconv.FormatInt(int64(clock), 10))
		} else {
			log.Printf(responseMessage.Name + ": " + responseMessage.Message + ", timestamp: " + strconv.FormatInt(int64(clock), 10))
		}
	}
}

func SendMessage(c chat.ChatClient) {
	for {
		clock += 1
		inputMessage, _ := reader.ReadString('\n')
		inputMessage = strings.TrimSpace(inputMessage)
		message := chat.WrittenMessage{Name: name, Message: inputMessage, TimeStamp: clock}
		messageAck, err := c.SendMessage(context.Background(), &message)
		syncClock(messageAck.TimeStamp)
		clock += 1
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
