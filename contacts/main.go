package main

import (
	"fmt"
	pb "github.com/iamyxsh/grpc-chat/contacts/contactspb"
	"github.com/iamyxsh/grpc-chat/contacts/db"
	"github.com/iamyxsh/grpc-chat/contacts/services"
	"github.com/iamyxsh/grpc-chat/kafka"
	"google.golang.org/grpc"
	"log"
	"net"
)

func logging(msg string) {
	fmt.Println(msg)
}

var addr string = "0.0.0.0:5002"

func main() {
	go kafka.ReadMsg("USER_LOGIN", logging)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterContactsServiceServer(s, &services.Server{})

	pg := db.ReturnDB("5433")
	db.ExecSchema(pg, db.Schema)

	//go kafka.CreateTopics()

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
