package services

import (
	"io"
	"log"
	"time"

	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
	"github.com/iamyxsh/grpc-chat/chat/db"
	"github.com/iamyxsh/grpc-chat/chat/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Channel struct {
	User   string
	Stream pb.ChatService_BroadcastMessageServer
}

var chanMap = make(map[string]Channel)

func (c *ChatService) BroadcastMessage(stream pb.ChatService_BroadcastMessageServer) error {

	pg := db.ReturnDB("5432")

	headers, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Error(codes.Internal, "Error while reading the context")
	}
	token := headers.Get("auth")
	if len(token) == 0 {
		return status.Error(codes.Unauthenticated, "Expected authorization header")
	}
	number, ok := utils.VerifyJWT(token[0])

	for {
		log.Println("First", chanMap)
		msg, err := stream.Recv()

		if err != nil {
			delete(chanMap, number)
			if err == io.EOF {
				break
			} else {
				log.Printf("Error in receiving message from client :: %v", err)
				break
			}

		}

		channel := Channel{
			User:   msg.From,
			Stream: stream,
		}

		if msg.Message == "" {

			chanMap[msg.From] = channel
			m := db.Messages{}
			messages := m.GetUndelivered(pg, msg.From)
			for _, mes := range messages {
				t, err := time.Parse(time.RFC3339Nano, mes.Timestamp)
				if err != nil {
					log.Println(err)
					break
				}
				err = stream.Send(&pb.Message{
					From:      mes.Sender,
					To:        mes.Receiver,
					Message:   mes.Body,
					Timestamp: t.Unix(),
				})
				if err != nil {
					log.Println(err)
				}
				m.MarkDelivered(pg, msg.From)
			}

		} else {
 
			exist := chanMap[msg.From]

			if exist.User != "" {
				chanMap[msg.From] = channel
				exist = channel
			}

			to := chanMap[msg.To]

			if to.User == "" {
				msg := db.Messages{
					Sender:    msg.From,
					Receiver:  msg.To,
					Body:      msg.Message,
					Delivered: false,
				}
				msg.Create(pg)
			} else {
				msg.Timestamp = time.Now().Unix()
				err = to.Stream.Send(msg)
				m := db.Messages{
					Sender:    msg.From,
					Receiver:  msg.To,
					Body:      msg.Message,
					Delivered: true,
				}
				m.Create(pg)

				if err != nil {
					log.Println(err)
				}
			}

			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}
