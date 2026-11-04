package main

import (
	"log"
	"net"

	gamev1 "github.com/kupr666/Pocker-Web/gen/go/game/v1"
	v1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	gameClient := gamev1.NewGameServiceClient(conn)

	listener, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	v1.RegisterLobbyServiceServer(server, &LobbyServer{
		game: gameClient,
	})

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
