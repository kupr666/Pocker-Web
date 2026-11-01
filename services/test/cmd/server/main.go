package main

import (
	"context"
	"log"
	"net"

	testv1 "github.com/kupr666/Pocker-Web/gen/go/test/v1"
	"google.golang.org/grpc"
)

type greetingServer struct {
	testv1.UnimplementedGreetingServiceServer
}

func (s *greetingServer) SayHello(
	ctx context.Context,
	req *testv1.SayHelloRequest,
) (*testv1.SayHelloResponse, error) {
	return &testv1.SayHelloResponse{
		Greeting: "hi" + req.GetName(),
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	testv1.RegisterGreetingServiceServer(server, &greetingServer{})
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
