package services

import (
	pb "github.com/iamyxsh/grpc-chat/chat/chatpb"
	"io"
	"log"
)

type Channel struct {
	User   string
	Stream pb.ChatService_BroadcastMessageServer
}

var chanMap = make(map[string]Channel)

func (c *ChatService) BroadcastMessage(stream pb.ChatService_BroadcastMessageServer) error {

	for {

		msg, err := stream.Recv()

		if err != nil {
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
			log.Println(chanMap)
		}

		exist := chanMap[msg.From]

		if exist.User != "" {
			chanMap[msg.From] = channel
			exist = channel
		}

		to := chanMap[msg.To]

		if to.User != "" {
			err = to.Stream.Send(msg)
		} else {
			stream.Send(msg)
		}

		if err != nil {
			log.Println(err)

		}
		log.Println(chanMap)
	}
	return nil
}
