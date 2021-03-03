package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8000"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

type Pod struct {
	name string
	ip   string
}

var pod = &Pod{
	name: os.Getenv("PODNAME"),
	ip:   os.Getenv("PODIP"),
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	m := fmt.Sprintf("Hello %s ! \n Pod name is %s \n; Pod IP is %s \n;", in.GetName(), pod.name, pod.ip)
	log.Printf("Received: %v", m)
	return &pb.HelloReply{Message: m}, nil
}

func main() {
	log.Println("start!")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		log.Println("Key: ", pair[0])
		log.Println("Value: ", pair[1])
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
