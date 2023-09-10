package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpc-demo/common"
)

const (
	port = ":50051"
)

type server struct {
	pb.GreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	endTime := time.Now()
	log.Printf("%s", endTime)

	reply := &pb.HelloReply{
		ResponseTimestamp: timestamppb.Now(),
	}
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.MaxConcurrentStreams(1000),
		grpc.InitialWindowSize(65536),
		grpc.InitialConnWindowSize(65536),
	)
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("Server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
