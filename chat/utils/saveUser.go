package utils

import (
	"encoding/json"
	"fmt"
	"github.com/iamyxsh/grpc-chat/chat/db"
	"log"
)

func SaveUser(msg string) {
	fmt.Println(msg)
	user := db.User{}
	err := json.Unmarshal([]byte(msg), user)

	if err != nil {
		log.Fatalln(err)
	}

	pg := db.ReturnDB("5432")
	user.Create(pg)
}
