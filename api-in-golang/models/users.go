package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(db *sql.DB, user User) (id int, err error) {
	stmt, errP := db.Prepare("INSERT INTO users(email,password) VALUES(?, ?)")
	result, errE := stmt.Exec(user.Email, user.Password)
	if errP != nil || errE != nil {
		log.Fatal("Insert user fail - executing query:", err)
		return 0, err
	}
	defer stmt.Close()
	id64, _ := result.LastInsertId()
	id = int(id64)
	return
}
