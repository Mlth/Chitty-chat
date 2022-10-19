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
	chat.UnimplementedChatServer
}

// Might give two users the same id
func (s *ChittyChatServer) JoinServer(ctx context.Context, in *chat.WrittenMessage) (*chat.WrittenMessage, error) {
	log.Printf("%s joined server. They have been assigned id: %d", in.Name, idCount)
	idCount++
	return &chat.WrittenMessage{Name: in.Name, Message: " ", TimeStamp: " ", Id: idCount - 1}, nil
}

func (s *ChittyChatServer) LeaveServer(ctx context.Context, in *chat.WrittenMessage) (*chat.WrittenMessage, error) {
	log.Printf("%s left the server", in.Name)
	return &chat.WrittenMessage{Name: in.Name, Message: " ", TimeStamp: " ", Id: in.Id}, nil
}

func (s *ChittyChatServer) SendMessage(ctx context.Context, in *chat.WrittenMessage) (*chat.WrittenMessage, error) {

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
