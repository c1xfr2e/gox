package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/bluecover/goexpr/grpc/hello"
	empty "github.com/golang/protobuf/ptypes/empty"
)

const (
	address = "localhost:50051"
)

func tryHelloWorld(cli pb.HelloServiceClient) {
	req := &empty.Empty{}
	resp, err := cli.HelloWorld(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not invoke Hello: %v", err)
	}
	fmt.Printf("HelloWorld response: %s\n", resp.Message)
}

func tryHello(cli pb.HelloServiceClient, msg string) {
	req := &pb.HelloRequest{Name: msg}
	resp, err := cli.Hello(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not invoke Hello: %v", err)
	}
	fmt.Printf("Hello %s, response: %s\n", msg, resp.Message)
}

func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a new HelloServiceClient
	cli := pb.NewHelloServiceClient(conn)
	tryHelloWorld(cli)
	tryHello(cli, "sqrt(zhaohao)")
}
