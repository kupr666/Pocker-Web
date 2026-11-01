package main

import (
	"context"
	"log"
	"time"

	testv1 "github.com/kupr666/Pocker-Web/gen/go/test/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		log.Fatal(err)
	}
	defer conn.Close()
	client := testv1.NewGreetingServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.SayHello(ctx, &testv1.SayHelloRequest{
		Name: "John Doe",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.GetGreeting())
}
