package db

import "github.com/jmoiron/sqlx"

type User struct {
	Number string `json:"number" db:"number"`
}

func (u User) Create(db *sqlx.DB) {
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO contacts (number) VALUES ($1)", u.Number)

	tx.Commit()
}

func (u User) GetByNumber(db *sqlx.DB) error {
	tx := db.MustBegin()

	user := User{}

	return tx.Get(&user, "SELECT * from users WHERE number=$1", u.Number)
}
