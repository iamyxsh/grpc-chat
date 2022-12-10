package main

import (
	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
	"github.com/iamyxsh/grpc-chat/chat/db"
	"github.com/iamyxsh/grpc-chat/chat/services"
	"github.com/iamyxsh/grpc-chat/chat/utils"
	"github.com/iamyxsh/grpc-chat/kafka"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:5003"

func main() {
	go kafka.ReadMsg("USER_LOGIN", utils.SaveUser)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &services.ChatService{})

	pg := db.ReturnDB("5432")
	db.ExecSchema(pg, db.Schema)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
