package controller

import (
	"fmt"
	"database/sql"
	"net/http"
	"github.com/thepiyush-303/rett/db"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(dbconn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		email := r.FormValue("email")
		password := r.FormValue("password")
		// username := r.FormValue("username")

		if email == "" || password == "" {
			http.Error(w, "missing fields", http.StatusBadRequest)
			return
		}

		_, err := db.checkUserByEmail(dbconn, email)

		if err == nil{
			http.Error(w, "user already exist", http.StatusBadRequest)
			return 
		}
		
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

		if err != nil{
			http.Error(w, "error hashing pass", http.StatusInternalServerError)
			return
		}

		password = string(hashedPassword)
		user := db.User{
			Email: email,
			Password: string(hashedPassword),
		}

		db.insertUser(dbconn, user)
		fmt.Fprintf(w, "user registered")
	}
}