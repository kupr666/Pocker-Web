package main

import (
	"log"
	"net/http"

	tablesv1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
	"github.com/kupr666/Pocker-Web/services/gateway/internal/interfaces/http/handlers/lobby"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// sens grpc in Lobby
	lobbyClient := tablesv1.NewLobbyServiceClient(conn)
	// needed to accept http request and call client
	gatewayHandler := lobby.NewGatewayServer(lobbyClient)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/tables", gatewayHandler.CreateTables)
	log.Println("Gateway listening at 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
