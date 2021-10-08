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

func OpenDatabase(source string) *sql.DB {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		log.Fatal("OpenDatabase(models) - fail to open database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("OpenDatabase(models) - fail to ping database:", err)
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

// InitializeDB with 'production' or 'test'
func InitializeDB(state string) *sql.DB {
	switch state {
	case "production":
		createStorageFolder()
		db := OpenDatabase(config.Database)
		//defer db.Close()
		startDatabase(db)
		return db
	case "test":
		db := OpenDatabase("file::memory:?cache=shared")
		//defer db.Close()
		startDatabase(db)
		return db
	default:
		log.Fatal("InitializeDB: no state provide")
		return nil
	}
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
