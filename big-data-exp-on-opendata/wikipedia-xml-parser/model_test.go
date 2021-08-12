package main

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_insertDatabase(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	defer db.Close()
	startDatabase(db)

	firstInsert := insertDatabase(db, 8, "title1", "https://url.test", "abstract", 2)
	secondInsert := insertDatabase(db, 7, "title2", "https://url.test", "abstract", 2)
	falseInsert := insertDatabase(db, 7, "title2", "https://url.test", "abstract", 2)
	if !firstInsert && !secondInsert && falseInsert {
		t.Errorf("insertDatabase tests fail")
	}

	os.Remove("./test.db")
}
