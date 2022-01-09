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

<<<<<<< HEAD
// Implementation of interface
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
func (server *UserNameServer) GetUserByName(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved Name: %v", un.GetName())

	const (
		address = "localhost:9001"
	)

<<<<<<< HEAD
	// Establishing connection at 9000
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect. Error: %v", err)
	}
	defer conn.Close()

<<<<<<< HEAD
	// Creating a new client of logic section
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

<<<<<<< HEAD
	// Collecting response from logic server
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	r, err := c.GetUser(ctx1, &pb1.Username{Name: un.GetName()})
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	// Returning data to client
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	roll, err := strconv.ParseInt(r.Roll, 10, 64)
	return &pb.User{Name: r.Name, Class: r.Class, Roll: roll}, err
}

func main() {
<<<<<<< HEAD

	// Listening to calls on port 9000
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. Error: %v", err)
	}
<<<<<<< HEAD

	// Creating a new server
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	server := grpc.NewServer()
	pb.RegisterUserNameServer(server, &UserNameServer{})
	log.Printf("Server listening at: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
