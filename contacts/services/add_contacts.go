package services

import (
	"context"

	pb "github.com/iamyxsh/grpc-chat/contacts/contactspb"
	"github.com/iamyxsh/grpc-chat/contacts/db"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (*Server) AddContacts(ctx context.Context, req *pb.ContactsRequest) (*pb.ContactsResponse, error) {
	pg := db.ReturnDB("5433")

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.ContactsResponse{}, status.Error(1, "something went wrong")
	}

	number := headers.Get("number")[0]

	for _, contact := range req.Number {

		contact := db.Contact{
			Contact: contact,
			Number:  number,
		}

		contact.Create(pg)
	}

	return &pb.ContactsResponse{Msg: ""}, nil
}
