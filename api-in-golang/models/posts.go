package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterPost(db *sql.DB, title string, datas string, idUser int) (id int, err error) {
	created := fmt.Sprintln(time.Now())

	insert, err := db.Begin()
	if err != nil {
		log.Fatal("Insert fail - opening database:", err)
		return 0, err
	}
	stmt, err := insert.Prepare("INSERT INTO posts(title, datas, created, idUser) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Insert fail - preparing query:", err)
		return 0, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(title, datas, created, idUser)
	if err != nil {
		log.Println("Insert fail - executing query:", err)
		return 0, err
	}
	insert.Commit()
	return
}
