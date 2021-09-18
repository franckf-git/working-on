package models

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"lite-api-crud/config"

	_ "github.com/mattn/go-sqlite3"
)

func createStorageFolder() {
	var folder string = strings.Split(config.Database, "/")[1]
	os.Mkdir(folder, 0755)
}

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", config.Database)
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	return db
}

func startDatabase(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		datas TEXT NOT NULL,
		created TEXT NOT NULL,
		idUser INTEGER NOT NULL,
		FOREIGN KEY(idUser) REFERENCES users(id)
		);
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user TEXT NOT NULL
		);`
	var err error
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println("Error during creating tables:", err, sqlStmt)
	}
}

func InitializeDB() {
	createStorageFolder()
	db := OpenDatabase()
	defer db.Close()
	startDatabase(db)
}

func CleanTables(db *sql.DB) {
	db.Exec("DELETE FROM posts")
	db.Exec("DELETE FROM users")
	db.Exec("VACUUM")
	db.Exec("UPDATE sqlite_sequence SET seq =0")
}
