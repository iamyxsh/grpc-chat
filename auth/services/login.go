package services

import (
	"encoding/json"
	pb "github.com/iamyxsh/grpc-chat/auth/authpb"
	"github.com/iamyxsh/grpc-chat/auth/db"
	"github.com/iamyxsh/grpc-chat/auth/utils"
	"github.com/iamyxsh/grpc-chat/kafka"
	"golang.org/x/net/context"
	"log"
)

func (*Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	pg := db.ReturnDB("5432")

	user := db.User{
		Number: req.Number,
	}

	err := user.GetByNumber(pg)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return &pb.LoginResponse{}, err

		}

		user.Create(pg)
	}

	msg, _ := json.Marshal(user)
	kafka.ProduceMessage("USER_LOGIN", string(msg))

	token, err := utils.GenerateJWT(user.Number)
	if err != nil {
		log.Print(err)
		return &pb.LoginResponse{}, err
	}

	return &pb.LoginResponse{Msg: token}, nil
}
