package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib" //-> driver
)

func connectDB(db *sql.DB){
	db, err := sql.Open(
		"pgx",
		"postgres://postgres:heythisismypassword@localhost:5432/authdb",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil{
		log.Fatal(err)
	}
	return db
}