package main

import (
	"log"
	"net"
	"strings"

	_ "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/bluecover/gox/grpc/hello"
)

const (
	port = ":50051"
)

type myserver struct{}

func (s *myserver) HelloWorld(context.Context, *empty.Empty) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello, world!"}, nil
}

func (s *myserver) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	r := &pb.HelloReply{
		Message: strings.ToUpper(req.Name),
	}
	return r, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC myserver
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &myserver{})
	s.Serve(lis)
}
