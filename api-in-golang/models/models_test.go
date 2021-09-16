package models

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test_RegisterPost(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	defer db.Close()
	StartDatabase(db)
	firstInsert, _ := RegisterPost(db, "title1", "datas1", 1)
	secondInsert, _ := RegisterPost(db, "title2", "datas2", 2)
	if firstInsert != 1 && secondInsert != 2 {
		t.Errorf("RegisterPost tests fail")
	}
	os.Remove("./test.db")
}
