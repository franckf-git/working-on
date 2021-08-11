package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// os.Remove("./enwiki-abstract.db")
	db, err := sql.Open("sqlite3", "./enwiki-abstract.db")
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	defer db.Close()

	startDatabase(db)
	insertDatabase(db, 8, "title", "https://url.test", "abstract", 2)
	insertDatabase(db, 9, "dex", "url", "abstract", 54)
	insertDatabase(db, 11, "TITLE", "https://url.fr", "abstract4", 462)
	output := selectAllDatabase(db)
	log.Println("Results:", output)
}

// startDatabase init database with tables doc and unknown
func startDatabase(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS doc(
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		abstract TEXT NOT NULL,
		links INTEGER NOT NULL
		);
	CREATE TABLE IF NOT EXISTS unknown(
		id INTEGER NOT NULL PRIMARY KEY,
		unknowntag TEXT NOT NULL,
		iddoc INTEGER NOT NULL,
		FOREIGN KEY(iddoc) REFERENCES doc(id)
		);`
	var err error
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Println("Error during creating tables:", err, sqlStmt)
	}
	log.Println("Database and tables ready.")
}

// insertDatabase insert values in databasee
func insertDatabase(db *sql.DB, id int, title string, url string, abstract string, links int) bool {
	insert, err := db.Begin()
	if err != nil {
		log.Fatal("Insert fail - opening database:", err)
		return false
	}
	stmt, err := insert.Prepare("INSERT INTO doc(id, title, url, abstract, links) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Insert fail - preparing query:", err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, title, url, abstract, links)
	if err != nil {
		log.Fatal("Insert fail - executing query:", err)
		return false
	}
	insert.Commit()
	return true
}

// selectAllDatabase return all lines in a slice of slices - VERY BAD DESIGN - TO redo
func selectAllDatabase(db *sql.DB) [][]string {
	result := make([][]string, 0)

	rows, err := db.Query("SELECT * FROM doc")
	if err != nil {
		log.Fatal("Select fail - executing query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var url string
		var abstract string
		var links int
		err = rows.Scan(&id, &title, &url, &abstract, &links)
		if err != nil {
			log.Fatal("Select fail - scanning values:", err)
		}
		currentRow := make([]string, 5)
		currentRow[0] = strconv.Itoa(id)
		currentRow[1] = title
		currentRow[2] = url
		currentRow[3] = abstract
		currentRow[4] = strconv.Itoa(links)
		result = append(result, currentRow)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Select fail - reading rows:", err)
	}
	return result
}
