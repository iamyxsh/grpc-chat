package services

import (
	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
}
