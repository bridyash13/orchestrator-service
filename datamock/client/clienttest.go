// THIS FILE IS ONLY FOR TESTING PURPOSE
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bridyash13/Mage/datamock/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserNameClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, err := client.GetMockUserData(ctx, &pb.Username{Name: "Yash"})
	if err != nil {
		log.Fatalf("Could not get mock user data. Error: %v", err)
	}
	log.Printf("User Details\nName: %s Age:%s Id: %d", r.GetName(), r.GetClass(), r.GetRoll())
}
