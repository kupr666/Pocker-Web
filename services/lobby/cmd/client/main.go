package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	v1 "github.com/kupr666/Pocker-Web/gen/go/tables/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("127.0.0.1:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := v1.NewLobbyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateTable(ctx, &v1.CreateTableRequest{
		OwnerUserId: "12",
		RequestId:   uuid.NewString(),
		Name:        "Cool Table",
		AccessMode:  v1.AccessMode_ACCESS_MODE_PUBLIC,
		Password:    "",
		Rules:       nil,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
