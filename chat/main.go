package main

import (
	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
	"github.com/iamyxsh/grpc-chat/chat/db"
	"github.com/iamyxsh/grpc-chat/chat/interceptor"
	"github.com/iamyxsh/grpc-chat/chat/services"
	"github.com/iamyxsh/grpc-chat/chat/utils"
	"github.com/iamyxsh/grpc-chat/kafka"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:5003"

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	opts := []grpc.ServerOption{}
	opts = append(opts, grpc.ChainStreamInterceptor(interceptor.StreamAuthInterceptor()))

	s := grpc.NewServer(opts...)
	pb.RegisterChatServiceServer(s, &services.ChatService{})

	pg := db.ReturnDB("5433")
	db.ExecSchema(pg, db.Schema)

	go kafka.CreateTopics()
	go kafka.ReadMsg("USER_LOGIN", utils.SaveUser)

	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
