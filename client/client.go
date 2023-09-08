package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	endTime := time.Now()
	log.Printf("%s", endTime)

	// Contact the server and print the response
	name := "World"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("Response: %s", r.Message)
}
