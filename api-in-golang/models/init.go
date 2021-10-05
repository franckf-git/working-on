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
		log.Fatal("OpenDatabase(models) - fail to open database:", err)
	}
	return db
}

func startDatabase(db *sql.DB) {
	createTablePosts := `
	CREATE TABLE IF NOT EXISTS posts(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		datas TEXT NOT NULL,
		created TEXT NOT NULL,
		idUser INTEGER NOT NULL,
		FOREIGN KEY(idUser) REFERENCES users(id)
		);`
	createTableUsers := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
		);`

	execDB(db, createTablePosts)
	execDB(db, createTableUsers)
}

func InitializeDB() {
	createStorageFolder()
	db := OpenDatabase()
	defer db.Close()
	startDatabase(db)
}

func CleanTables(db *sql.DB) {
	execDB(db, "DELETE FROM posts")
	execDB(db, "DELETE FROM users")
	execDB(db, "VACUUM")
	execDB(db, "UPDATE sqlite_sequence SET seq =0")
}

func execDB(db *sql.DB, request string) {
	stmt, err := db.Prepare(request)
	if err != nil {
		config.ErrorLogg(request, " - preparing query:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		config.ErrorLogg(request, " - creating tables:", err)
	}
}
