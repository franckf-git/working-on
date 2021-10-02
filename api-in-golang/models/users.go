package models

import (
	"database/sql"
	"fmt"
	"lite-api-crud/config"

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
		config.ErrorLogg("RegisterUser(models) - crypting password:", err)
		return 0, err
	}

	stmt, errP := db.Prepare("INSERT INTO users(email,password) VALUES(?, ?)")
	result, errE := stmt.Exec(user.Email, string(hash))
	if errP != nil || errE != nil {
		config.ErrorLogg("RegisterUser(models) - executing query:", err)
		return 0, err
	}
	defer stmt.Close()
	id64, _ := result.LastInsertId()
	id = int(id64)
	return
}

func (user *User) CheckExistingUser(db *sql.DB) (id int, err error) {
	stmt, errP := db.Prepare("SELECT * FROM users WHERE email=?")
	var email string
	var hashpassword string
	errQ := stmt.QueryRow(user.Email).Scan(&id, &email, &hashpassword)
	errC := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(user.Password))
	if errP != nil || errQ != nil || errC != nil || id == 0 {
		config.ErrorLogg("CheckExistingUser(models) - user and password:", errP, errQ, errC, id)
		return 0, fmt.Errorf("error checking user and password: %v, %v, %v, %v", errP, errQ, errC, id)
	}
	defer stmt.Close()
	return
}
