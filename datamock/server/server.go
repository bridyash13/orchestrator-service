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

<<<<<<< HEAD
// Implementation of interface
func (server *UserNameServer) GetMockUserData(ctx context.Context, un *pb.Username) (*pb.User, error) {
=======
func (s *UserNameServer) GetMockUserData(ctx context.Context, un *pb.Username) (*pb.User, error) {
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	log.Printf("Recieved: %v", un.GetName())
	if (len(un.GetName())) < 6 {
		return nil, errors.New("length is too short")
	}
	return &pb.User{Name: un.GetName(), Class: strconv.Itoa(len(un.GetName())), Roll: int64(len(un.GetName()) * 10)}, nil
}

func main() {
<<<<<<< HEAD

	// Listening to calls on port 10000
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. Error: %v", err)
	}
<<<<<<< HEAD

	// Creating new server
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to connect to server. Error: %v", err)
	}
<<<<<<< HEAD

=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
}
