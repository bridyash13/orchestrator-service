package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	pb "github.com/bridyash13/Mage/datamock/proto"
	"google.golang.org/grpc"
)

const (
	port = ":10000"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

// Implementation of interface
func (server *UserNameServer) GetMockUserData(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved: %v", un.GetName())
	if (len(un.GetName())) < 6 {
		return nil, errors.New("length is too short")
	}
	return &pb.User{Name: un.GetName(), Class: strconv.Itoa(len(un.GetName())), Roll: int64(len(un.GetName()) * 10)}, nil
}

func main() {

	// Listening to calls on port 10000
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. Error: %v", err)
	}

	// Creating new server
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to connect to server. Error: %v", err)
	}

}
