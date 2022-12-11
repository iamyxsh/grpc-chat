package services

import (
	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
	"github.com/iamyxsh/grpc-chat/chat/db"
	"log"
	"time"
)

func (c *ChatService) GetMessages(req *pb.GetMessagesRequest, stream pb.ChatService_GetMessagesServer) error {
	msg := db.Messages{
		Sender: req.User,
	}

	pg := db.ReturnDB("5432")

	messages := msg.GetByUser(pg, req.User)

	for _, message := range messages {
		t, err := time.Parse(time.RFC3339Nano, message.Timestamp)
		if err != nil {
			log.Println(err)
			break
		}

		go func() {
			resp := pb.Message{
				From:      message.Sender,
				To:        message.Receiver,
				Message:   message.Body,
				Timestamp: t.Unix(),
			}
			if err := stream.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		}()
	}

	return nil
}
