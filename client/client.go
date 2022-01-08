package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bridyash13/Mage/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection was not successful. Error: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	response, err := c.GetUserByName(ctx, &pb.Username{Name: "Yash"})
	if err != nil {
		log.Fatalf("Username not found. Error: %v", err)
	}
	log.Printf("Connection Successful! User Details are\n Name: %s Age: %s ID: %d", response.GetName(), response.GetClass(), response.GetRoll())
}
