package main

import (
	"log"
	pb "microservice/auth/proto/microservice/auth"
	service "microservice/auth/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":10001")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &service.ServiceServer{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
