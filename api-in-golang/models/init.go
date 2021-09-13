package models

import (
	"database/sql"
	"log"

	"lite-api-crud/config"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", config.Database)
	log.Println(config.Database)
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	return db
}

func StartDatabase(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		datas TEXT NOT NULL,
		created TEXT NOT NULL,
		idUser INTEGER NOT NULL,
		FOREIGN KEY(idUser) REFERENCES users(id)
		);
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY,
		user TEXT NOT NULL
		);`
	var err error
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println("Error during creating tables:", err, sqlStmt)
	}
}
