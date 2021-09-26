package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) RegisterUser(db *sql.DB) (id int, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error crypting password:", err)
		return 0, err
	}

	stmt, errP := db.Prepare("INSERT INTO users(email,password) VALUES(?, ?)")
	result, errE := stmt.Exec(user.Email, string(hash))
	if errP != nil || errE != nil {
		log.Println("Insert user fail - executing query:", err)
		return 0, err
	}
	defer stmt.Close()
	id64, _ := result.LastInsertId()
	id = int(id64)
	return
}
