package main

import (
	"context"
	"log"
	"sync"
	"time"

	pb "grpc-demo/common"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	address  = "localhost:50051"
	poolSize = 20
)

type ConnectionPool struct {
	pool *sync.Pool
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		pool: &sync.Pool{
			New: func() interface{} {
				conn, err := grpc.Dial(address, grpc.WithInsecure())
				if err != nil {
					log.Fatalf("Failed to create connection: %v", err)
				}
				return conn
			},
		},
	}
}

func (p *ConnectionPool) Get() *grpc.ClientConn {
	return p.pool.Get().(*grpc.ClientConn)
}

func (p *ConnectionPool) Put(conn *grpc.ClientConn) {
	p.pool.Put(conn)
}

func main() {
	// Create a connection pool
	pool := NewConnectionPool()
	latencies := make([]time.Duration, 0)

	// Make 100 requests to the server
	for i := 0; i < 100; i++ {
		// Get a connection from the pool
		conn := pool.Get()
		defer conn.Close() // Close the connection when done with it

		// Create a gRPC client using the connection
		c := pb.NewGreeterClient(conn)

		request := &pb.HelloRequest{
			RequestTimestamp: timestamppb.Now(),
		}

		// Send the request to the server
		response, err := c.SayHello(context.Background(), request)
		if err != nil {
			log.Fatalf("Failed to call SayHello: %v", err)
		}

		// Calculate and log the sending latency
		latency := response.ResponseTimestamp.AsTime().Sub(request.RequestTimestamp.AsTime())
		log.Printf("Sending Latency: %v", latency)

		// Store the latency in the slice
		latencies = append(latencies, latency)
	}

	// Calculate the average latency
	averageLatency := calculateAverageLatency(latencies)
	log.Printf("Average Latency: %v", averageLatency)

}

func calculateAverageLatency(latencies []time.Duration) time.Duration {
	if len(latencies) == 0 {
		return 0
	}

	var totalLatency time.Duration
	for _, latency := range latencies {
		totalLatency += latency
	}

	return totalLatency / time.Duration(len(latencies))
}
