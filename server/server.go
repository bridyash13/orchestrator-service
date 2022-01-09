package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	pb1 "github.com/bridyash13/Mage/logic/proto"
	pb "github.com/bridyash13/Mage/proto"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

// Implementation of interface
func (server *UserNameServer) GetUserByName(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved Name: %v", un.GetName())

	const (
		address = "localhost:9001"
	)

	// Establishing connection at 9000
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect. Error: %v", err)
	}
	defer conn.Close()

	// Creating a new client of logic section
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Collecting response from logic server
	r, err := c.GetUser(ctx1, &pb1.Username{Name: un.GetName()})
	if err != nil {
		return nil, err
	}

	// Returning data to client
	roll, err := strconv.ParseInt(r.Roll, 10, 64)
	return &pb.User{Name: r.Name, Class: r.Class, Roll: roll}, err
}

func main() {
	// Listening to calls on port 9000
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. Error: %v", err)
	}

	// Creating a new server
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
