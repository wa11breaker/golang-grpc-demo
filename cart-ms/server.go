package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	pb "cart-ms/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	serverPort         = ":50051"
)

// cartServiceServer implements the CartServiceServer interface
type cartServiceServer struct {
	pb.UnimplementedCartServiceServer
}

// GetCart implements the GetCart RPC method
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
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	s := grpc.NewServer()

	// Register the CartServiceServer
	pb.RegisterCartServiceServer(s, &cartServiceServer{})

	fmt.Println("Cart ms is running on :50051")

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
}
