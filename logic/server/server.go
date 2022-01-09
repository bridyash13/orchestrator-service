package main

import (
	"context"
	"log"
	"net"
	"strconv"
	"time"

	pb1 "github.com/bridyash13/Mage/datamock/proto"
	pb "github.com/bridyash13/Mage/logic/proto"
<<<<<<< HEAD

=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

type UserNameServer struct {
	pb.UnimplementedUserNameServer
}

<<<<<<< HEAD
// Function to fetch user
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
func (s *UserNameServer) GetUser(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved: %v", un.GetName())

	const (
		address = "localhost:10000"
	)

<<<<<<< HEAD
	// Establishing a connection on port 9001
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
<<<<<<< HEAD

	// Creating a new client of datamock section
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

<<<<<<< HEAD
	// Collecting response from datamock server
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	r, err := c.GetMockUserData(ctx1, &pb1.Username{Name: un.GetName()})
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	// Returning data to server
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	roll := strconv.Itoa(int(r.Roll))
	return &pb.User{Name: r.Name, Class: r.Class, Roll: roll}, err
}

func main() {
<<<<<<< HEAD

	// Listening to calls on port 9001
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
<<<<<<< HEAD

	// Creating new server
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
=======
	s := grpc.NewServer()
	pb.RegisterUserNameServer(s, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
		log.Fatalf("Failed to serve: %v", err)
	}
}
