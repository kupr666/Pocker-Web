package main

import (
	"log"
	"net"

	v1 "github.com/kupr666/Pocker-Web/gen/go/game/v1"
	"github.com/kupr666/Pocker-Web/services/game/internal/interfaces/gamegrpc"
	"google.golang.org/grpc"
)

func main() {
	// quick implementation to get minimal version of get table
	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	// регистрирует в gRPC-сервере методы GameService, объявленные в .proto.
	// Она не регистрирует все произвольные методы структуры GameServer.
	v1.RegisterGameServiceServer(server, gamegrpc.NewServer())

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
