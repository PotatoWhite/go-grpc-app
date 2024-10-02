package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "go-grpc-app/pkg/server/generated"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) Greet(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	return &pb.GreetingResponse{Greeting: "hello"}, nil
}

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})

	fmt.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
