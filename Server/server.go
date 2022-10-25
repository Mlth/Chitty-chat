package main

import (
	"context"
	"log"
	"net"

	chat "github.com/Mlth/Chitty-chat/proto"
	"google.golang.org/grpc"
)

var idCount int32 = 0

type ChittyChatServer struct {
	chat.ChatServer
}

var streams = make([]chat.Chat_JoinServerServer, 0)

// Might give two users the same id
func (s *ChittyChatServer) JoinServer(in *chat.WrittenMessage, stream chat.Chat_JoinServerServer) error {
	log.Printf("%s joined server. They have been assigned id: %d", in.Name, idCount)
	//idCount++

	streams = append(streams, stream)
	broadcastToAll(&chat.WrittenMessage{Message: in.Name + " has joined the server"})

	for {

	}
}

func broadcastToAll(in *chat.WrittenMessage) {
	for i := 0; i < len(streams); i++ {
		streams[i].Send(in)
	}
}

func (s *ChittyChatServer) SendMessage(ctx context.Context, in *chat.WrittenMessage) (*chat.EmptyMessage, error) {
	broadcastToAll(in)
	return &chat.EmptyMessage{}, nil
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
