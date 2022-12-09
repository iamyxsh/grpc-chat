package db

import (
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/lib/pq"
)

func ReturnDB(port string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@contacts-db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func ExecSchema(db *sqlx.DB, schema string) {
	db.MustExec(schema)
}
