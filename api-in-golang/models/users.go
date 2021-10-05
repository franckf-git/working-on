package models

import (
	"database/sql"
	"lite-api-crud/config"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) RegisterUser(db *sql.DB) (id int, err error) {
	stmt, err := db.Prepare("INSERT INTO users(email,password) VALUES(?, ?)")
	if err != nil {
		config.ErrorLogg("RegisterUser(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		config.ErrorLogg("RegisterUser(models) - crypting password:", err)
		return
	}

	result, err := stmt.Exec(user.Email, string(hash))
	if err != nil {
		config.ErrorLogg("RegisterUser(models) - executing query:", err)
		return
	}

	idReturn, _ := result.LastInsertId()
	id = int(idReturn)
	return
}

func (user *User) CheckExistingUser(db *sql.DB) (id int, err error) {
	stmt, err := db.Prepare("SELECT * FROM users WHERE email=?")
	if err != nil {
		config.ErrorLogg("CheckExistingUser(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	var email string
	var hashpassword string
	err = stmt.QueryRow(user.Email).Scan(&id, &email, &hashpassword)
	if err != nil {
		config.ErrorLogg("CheckExistingUser(models) - reading rows:", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(user.Password))
	if err != nil {
		config.ErrorLogg("CheckExistingUser(models) - compare hash password:", err)
		return
	}
	return
}
