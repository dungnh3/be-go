package main

import (
	pb "be-go/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	clientConn, err := grpc.Dial("localhost:10443", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err while dial %v", err)
	}
	defer clientConn.Close()

	client := pb.NewDemoServiceClient(clientConn)
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
