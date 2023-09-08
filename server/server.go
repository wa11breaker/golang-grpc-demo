package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "grpc-demo/grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.GreeterServer
}

// SayHello implements the SayHello gRPC method
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	endTime := time.Now()
	log.Printf("%s", endTime)

	reply := &pb.HelloReply{Message: "Hello, " + in.Name + "!"}
	return reply, nil
}

// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
// }

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("Server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
