package services

import (
	"context"
	pb "github.com/iamyxsh/grpc-chat/contacts/contactspb"
	"github.com/iamyxsh/grpc-chat/contacts/db"
)

func (*Server) AddContacts(ctx context.Context, req *pb.ContactsRequest) (*pb.ContactsResponse, error) {
	pg := db.ReturnDB("5433")

	for _, num := range req.Number {
		tx := pg.
	}

	return &pb.ContactsResponse{Msg: ""}, nil
}
