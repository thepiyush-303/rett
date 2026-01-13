package db

import (
	"database/sql"
	"log"
)


func CheckUserByEmail(db *sql.DB, email string) (User, error){
	query := `SELECT email, password FROM users WHERE email = $1`

	var user User
	
	err := db.QueryRow(query, email).Scan(&user.Email, &user.Password)

	if err != nil{
		return user, err
	}
	return user, nil
}

func InsertUser(db *sql.DB, user User) int {
	query := `INSERT INTO users ( 
	name, email) VALUES ($1, $2) RETURNING id`

	var pk int

	err := db.QueryRow(query, user.Email, user.Password).Scan(&pk)

	if err != nil{
		log.Fatal(err)
	}
	return pk
}