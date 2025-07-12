package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/proto/gen"
	"log"
	"net"
)

type server struct {
	mainapipb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, request *mainapipb.AddRequest) (*mainapipb.AddResponse, error) {
	return &mainapipb.AddResponse{
		Sum: request.A + request.B,
	}, nil
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Error: Failed to listen!")
	}

	grpcServer := grpc.NewServer()

	mainapipb.RegisterCalculateServer(grpcServer, &server{})

	fmt.Println("Server running on port : ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Println("Error: Failed to serve!")
	}
}
