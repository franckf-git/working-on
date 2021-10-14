package models

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

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
	case "test":
		log.Printf("config.State init: %#+v\n", config.State)
		db := OpenDatabase("file::memory:?cache=shared")
		startDatabase(db)
		return db
	default:
		createStorageFolder()
		db := OpenDatabase(config.Database)
		startDatabase(db)
		backupDatabase()
		migrateDatabase()
		return db
	}
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

func migrateDatabase() {
	fileInfo, err := ioutil.ReadDir("./models/migrate/")
	if err != nil {
		config.ErrorLogg("reading content of migrate folder", err)
	}
	for _, file := range fileInfo {
		args := "sqlite3 " + config.DatabaseFile + " < " + "./models/migrate/" + file.Name()
		out, err := exec.Command("bash", "-c", args).CombinedOutput()
		if err != nil {
			config.ErrorLogg("migrating", file.Name(), string(out), err)
		}
	}
}

func backupDatabase() {
	timer := string(time.Now().Format("2006-01-02_15:04:05"))
	args := "sqlite3 " + config.DatabaseFile + " .dump " + " > " + "./models/backup/" + timer + ".sql"
	out, err := exec.Command("bash", "-c", args).CombinedOutput()
	if err != nil {
		config.ErrorLogg("backup", timer, string(out), err)
	}
}
