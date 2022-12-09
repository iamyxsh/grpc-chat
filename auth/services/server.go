package services

import pb "github.com/iamyxsh/grpc-chat/auth/authpb"

type Server struct {
	pb.AuthServiceServer
}
