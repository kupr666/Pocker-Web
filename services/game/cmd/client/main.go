package main

import (
	"context"
	"log"
	"time"

	v1 "github.com/kupr666/Pocker-Web/gen/go/game/v1"
	tablev1 "github.com/kupr666/Pocker-Web/gen/go/table/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		log.Fatal(err)
	}
	defer conn.Close()
	client := v1.NewGameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.CreateTableRuntime(ctx, &v1.CreateTableRuntimeRequest{
		TableId: "demo-table-id",
		Rules: &tablev1.TableRules{
			MaxSeats:           6,
			StartingStack:      1000,
			SmallBlind:         5,
			BigBlind:           10,
			TurnTimeoutSeconds: 30,
			ReentryMode:        tablev1.ReentryMode_REENTRY_MODE_PRESERVE_STACK,
			RebuyMode:          tablev1.RebuyMode_REBUY_MODE_PLAYER_REQUEST,
			RebuyLimit:         3,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.GetTableId())
}
