package main

import (
	"context"
	"log"
	"net"
	"strconv"

	chat "github.com/Mlth/Chitty-chat/proto"
	"google.golang.org/grpc"
)

type ChittyChatServer struct {
	chat.ChatServer
}

var streams = make([]chat.Chat_JoinServerServer, 0)
var clock int32 = 0

// Might give two users the same id
func (s *ChittyChatServer) JoinServer(in *chat.WrittenMessage, stream chat.Chat_JoinServerServer) error {
	log.Printf(in.Name + " has joined the server - timestamp: " + strconv.FormatInt(int64(in.TimeStamp), 10))

	streams = append(streams, stream)

	syncClock(in.TimeStamp)
	clock += 1
	broadcastToAll(&chat.WrittenMessage{Message: in.Name + " has joined the server", TimeStamp: clock})

	var streamClosed bool = false
	for {
		select {
		case <-stream.Context().Done():
			clock += 1
			log.Printf("%s left the server", in.Name)
			broadcastToAll(&chat.WrittenMessage{Message: in.Name + " has left the server", TimeStamp: clock})
			streamClosed = true
		}
		if streamClosed {
			break
		}
	}
	return nil
}

func broadcastToAll(in *chat.WrittenMessage) {
	log.Printf("sending message to clients - timestamp: " + strconv.FormatInt(int64(clock), 10))
	for i := 0; i < len(streams); i++ {
		streams[i].Send(in)
	}
}

func (s *ChittyChatServer) SendMessage(ctx context.Context, in *chat.WrittenMessage) (*chat.EmptyMessage, error) {
	log.Printf("recieving message from client: " + in.Name + " - timestamp: " + strconv.FormatInt(int64(in.TimeStamp), 10))
	syncClock(in.TimeStamp)
	clock += 1
	broadcastToAll(in)
	return &chat.EmptyMessage{TimeStamp: clock}, nil
}

func main() {
	// Create listener tcp on port 9080
	list, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen on port 9080: %v", err)
	}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServer(grpcServer, &ChittyChatServer{})
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

func syncClock(timestamp int32) {

	if clock < timestamp {
		clock = timestamp
	}
}
