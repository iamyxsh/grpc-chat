package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Contact struct {
	Number  string `json:"number" db:"number"`
	Contact string `json:"contact" db:"contact"`
}

func (c Contact) Create(db *sqlx.DB) {
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO contacts (number, contact) VALUES ($1, $2) ON CONFLICT DO NOTHING;", c.Number, c.Contact)

	tx.Commit()
}

func (c Contact) GetByNumber(db *sqlx.DB) ([]Contact, error) {
	tx := db.MustBegin()

	contacts := []Contact{}

	err := tx.Select(&contacts, "SELECT * FROM WHERE number = $1;", c.Number)

	return contacts, err
}

func (c Contact) Delete(db *sqlx.DB) {
	tx := db.MustBegin()

	fmt.Println(c.Number)

	tx.MustExec("DELETE FROM contacts WHERE number = $1 AND contact = $2;", c.Number, c.Contact)

	tx.Commit()
}
