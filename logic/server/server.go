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

func (s *UserNameServer) GetUser(ctx context.Context, un *pb.Username) (*pb.User, error) {
	log.Printf("Recieved: %v", un.GetName())

	const (
		address = "localhost:10000"
	)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := pb1.NewUserNameClient(conn)

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	r, err := c.GetMockUserData(ctx1, &pb1.Username{Name: un.GetName()})
	if err != nil {
		return nil, err
	}

	roll := strconv.Itoa(int(r.Roll))
	return &pb.User{Name: r.Name, Class: r.Class, Roll: roll}, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserNameServer(s, &UserNameServer{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
