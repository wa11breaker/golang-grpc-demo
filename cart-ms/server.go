package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc-demo/cart-ms/proto"
	"io/ioutil"
	"log"
	"net"
	"time"
)

// Struct that implements the CartServiceServer interface
type cartServiceServer struct {
	pb.CartServiceServer
}

// Implement the GetCart RPC method
func (s *cartServiceServer) GetCart(ctx context.Context, empty *pb.Empty) (*pb.CartResponse, error) {
	startTime := time.Now()
	data, err := ioutil.ReadFile("data/data_10kb.json")
	if err != nil {
		log.Printf("Error reading data file: %v", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	var cart pb.CartResponse
	err = json.Unmarshal(data, &cart)
	if err != nil {
		log.Printf("Error unmarshaling JSON data: %v", err)
		return nil, status.Error(codes.Internal, "Internal Server Error")
	}

	elapsedTimeMs := time.Since(startTime).Milliseconds()
	log.Printf("Execution time: %d ms", elapsedTimeMs)

	return &cart, nil
}

func main() {
	// Create a gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()

	// Register the CartServiceServer
	pb.RegisterCartServiceServer(s, &cartServiceServer{})

	fmt.Println("gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
