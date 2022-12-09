package main

import (
	"github.com/iamyxsh/grpc-chat/auth/db"
	"github.com/iamyxsh/grpc-chat/kafka"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"

	pb "github.com/iamyxsh/grpc-chat/auth/authpb"
	"github.com/iamyxsh/grpc-chat/auth/services"
)

var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:5001"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &services.Server{})

	pg := db.ReturnDB("5432")
	db.ExecSchema(pg, db.Schema)

	go kafka.CreateTopics()

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
