package main

import (
	"database/sql"
	"log"
	"os"
	"reflect"
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

func Test_selectAllDatabase(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	defer db.Close()
	startDatabase(db)

	insertDatabase(db, 8, "title1", "https://url.test", "abstract", 2)
	insertDatabase(db, 7, "title2", "https://url.test", "abstract", 2)
	want := [][]string{{"7", "title2", "https://url.test", "abstract", "2"}, {"8", "title1", "https://url.test", "abstract", "2"}}
	got := selectAllDatabase(db)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("selectAllDatabasetests fail: want %v - got %v", want, got)
	}

	os.Remove("./test.db")
}
