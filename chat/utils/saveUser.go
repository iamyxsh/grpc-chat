package utils

import (
	"encoding/json"
	"log"

	"github.com/iamyxsh/grpc-chat/chat/db"
)

func SaveUser(msg string) {
	user := db.User{}
	err := json.Unmarshal([]byte(msg), &user)

	if err != nil {
		log.Fatalln(err)
	}

	pg := db.ReturnDB("5432")
	user.Create(pg)
}
