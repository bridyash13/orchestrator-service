package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/bridyash13/Mage/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {

	// Taking input from the user
	fmt.Println("Enter Your First Name: ")
	var name string
	fmt.Scanln(&name)

	// Establishing a connection at port 9000
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection was not successful. Error: %v", err)
	}
	defer conn.Close()


	// Creating a new client
	c := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()


	// Calling a remote procedure
	response, err := c.GetUserByName(ctx, &pb.Username{Name: name})
	if err != nil {
		log.Fatalf("Username not found. Error: %v", err)
	}
	log.Printf("Connection Successful! User Details are\n Name: %s Age: %s ID: %d", response.GetName(), response.GetClass(), response.GetRoll())
}
