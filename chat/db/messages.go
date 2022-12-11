package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Messages struct {
	Sender    string `json:"sender" db:"sender"`
	Receiver  string `json:"receiver" db:"receiver"`
	Body      string `json:"body" db:"body"`
	Timestamp string `json:"timestamp" db:"timestamp"`
	Delivered bool   `json:"delivered" db:"delivered"`
}

func (m Messages) Create(db *sqlx.DB) {
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO messages (sender, receiver, body, delivered) VALUES ($1, $2, $3, $4)", m.Sender, m.Receiver, m.Body, m.Delivered)

	tx.Commit()
}

func (m Messages) GetByUser(db *sqlx.DB, user string) []Messages {
	tx := db.MustBegin()

	messages := []Messages{}

	err := tx.Select(&messages, "SELECT * from messages WHERE sender = $1 OR receiver = $1", user)

	if err != nil {
		log.Println(err)
	}

	return messages
}

func (m Messages) GetUndelivered(db *sqlx.DB, user string) []Messages {
	tx := db.MustBegin()

	messages := []Messages{}

	err := tx.Select(&messages, "SELECT * from messages WHERE receiver = $1 AND delivered = false", user)

	if err != nil {
		log.Println(err)
	}

	return messages
}

func (m Messages) MarkDelivered(db *sqlx.DB, user string) {
	tx := db.MustBegin()

	tx.Exec("UPDATE messages SET delivered = true WHERE receiver = $1 AND delivered = false", user)

	tx.Commit()
}
