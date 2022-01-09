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

<<<<<<< HEAD
	// Taking input from the user
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	fmt.Println("Enter Your First Name: ")
	var name string
	fmt.Scanln(&name)

<<<<<<< HEAD
	// Establishing a connection at port 9000
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection was not successful. Error: %v", err)
	}
	defer conn.Close()

<<<<<<< HEAD
	// Creating a new client
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	c := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

<<<<<<< HEAD
	// Calling a remote procedure
=======
>>>>>>> 02ba428953dd1c9662d31b629d62e00ea2eadbe4
	response, err := c.GetUserByName(ctx, &pb.Username{Name: name})
	if err != nil {
		log.Fatalf("Username not found. Error: %v", err)
	}
	log.Printf("Connection Successful! User Details are\n Name: %s Age: %s ID: %d", response.GetName(), response.GetClass(), response.GetRoll())
}
