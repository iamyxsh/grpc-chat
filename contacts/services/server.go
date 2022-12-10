package services

import (
	pb "github.com/iamyxsh/grpc-chat/contacts/contactspb"
)

type Server struct {
	pb.UnimplementedContactsServiceServer
}
