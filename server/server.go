package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/bridyash13/orchestrator-service/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type UserServer struct {
	pb.UnimplementedUserNameServer
}

func (s *UserServer) GetUserByNamme(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved Name: %v", un.GetName())
	return nil, errors.New("not implemented yet. Yash will implement me")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. Error: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserServer{})
	log.Printf("Server listening at: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
