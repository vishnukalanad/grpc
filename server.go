package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/proto/gen"
	"grpc/proto/gen/farewell"
	"log"
	"net"
)

type server struct {
	mainapipb.UnimplementedCalculateServer
	mainapipb.UnimplementedBidFarewellServer
	//farewellpb.UnimplementedFarewellServer
}

type serverGreeter struct {
	mainapipb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, request *mainapipb.AddRequest) (*mainapipb.AddResponse, error) {
	sum := request.A + request.B
	log.Println("SUM : ", sum)
	return &mainapipb.AddResponse{
		Sum: request.A + request.B,
	}, nil
}

func (s *server) Greeter(ctx context.Context, request *mainapipb.HelloRequest) (*mainapipb.HelloResponse, error) {
	return &mainapipb.HelloResponse{
		Message: "Hello " + request.Name,
	}, nil
}

func (s *server) BidGoodBye(ctx context.Context, request *farewellpb.GoodByeRequest) (*farewellpb.GoodByeResponse, error) {
	log.Println("Greet server request !")
	return &farewellpb.GoodByeResponse{
		Message: "Good bye" + request.Name,
	}, nil
}

func main() {

	cert := "cert.pem"
	key := "key.pem"

	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error: Failed to listen!")
		return
	}

	cred, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		log.Fatal("Error: Failed to load credentials!")
		return
	}
	grpcServer := grpc.NewServer(grpc.Creds(cred))

	mainapipb.RegisterCalculateServer(grpcServer, &server{})
	mainapipb.RegisterGreeterServer(grpcServer, &server{})
	//farewellpb.RegisterFarewellServer(grpcServer, &server{})
	mainapipb.RegisterBidFarewellServer(grpcServer, &server{})

	fmt.Println("Server running on port : ", port)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Error: Failed to serve!")
		return
	}
}
