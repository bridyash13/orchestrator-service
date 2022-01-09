package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	pb1 "github.com/bridyash13/Mage/datamock/proto"
	pb "github.com/bridyash13/Mage/logic/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

// Function to fetch user
func (s *UserNameServer) GetUser(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved: %v", un.GetName())

	const (
		address = "localhost:10000"
	)

	// Establishing a connection on port 9001
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// Creating a new client of datamock section
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Collecting response from datamock server
	r, err := c.GetMockUserData(ctx1, &pb1.Username{Name: un.GetName()})
	if err != nil {
		return nil, err
	}

	// Returning data to server
	roll := strconv.Itoa(int(r.Roll))
	return &pb.User{Name: r.Name, Class: r.Class, Roll: roll}, err
}

func main() {

	// Listening to calls on port 9001
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	// Creating new server
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
