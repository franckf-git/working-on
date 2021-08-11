package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./enwiki-abstract.db")
	db, err := sql.Open("sqlite3", "./enwiki-abstract.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	startDatabase(db)

	insert, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := insert.Prepare("insert into doc(id, title, url, abstract, links) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, "title", "https://url.test", "abstract", i)
		if err != nil {
			log.Fatal(err)
		}
	}
	insert.Commit()

	rows, err := db.Query("select id, title from doc")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func startDatabase(db *sql.DB) {

	sqlStmt := `
	CREATE TABLE doc(
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		abstract TEXT NOT NULL,
		links INTEGER NOT NULL
		);
	CREATE TABLE unknown(
		id INTEGER NOT NULL PRIMARY KEY,
		unknowntag TEXT NOT NULL,
		iddoc INTEGER NOT NULL,
		FOREIGN KEY(iddoc) REFERENCES doc(id)
		);
	`

	var err error
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	log.Println("database created")
}
