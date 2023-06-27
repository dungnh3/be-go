package main

import (
	pb "be-go/pb"
	context "context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
}

func (s Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "hello",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:10443")
	if err != nil {
		log.Fatalf("err while create listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterDemoServiceServer(s, &Server{})

	log.Println("calculator is running..")
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
