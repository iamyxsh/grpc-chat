package main

import (
	pb "github.com/iamyxsh/grpc-chat/contacts/contactspb"
	"github.com/iamyxsh/grpc-chat/contacts/db"
	"github.com/iamyxsh/grpc-chat/contacts/interceptor"
	"github.com/iamyxsh/grpc-chat/contacts/services"
	"github.com/iamyxsh/grpc-chat/contacts/utils"
	"github.com/iamyxsh/grpc-chat/kafka"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:5002"

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.ChainUnaryInterceptor(interceptor.CheckHeaderInterceptor()))

	s := grpc.NewServer(opts...)
	pb.RegisterContactsServiceServer(s, &services.Server{})

	pg := db.ReturnDB("5433")
	db.ExecSchema(pg, db.Schema)

	go kafka.CreateTopics()
	go kafka.ReadMsg("USER_LOGIN", utils.SaveUser)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
