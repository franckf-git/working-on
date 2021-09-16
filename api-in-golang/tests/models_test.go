package tests

import (
	"database/sql"
	"lite-api-crud/models"
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
	models.StartDatabase(db)
	firstInsert, _ := models.RegisterPost(db, "title1", "datas1", 1)
	secondInsert, _ := models.RegisterPost(db, "title2", "datas2", 2)
	if firstInsert != 1 && secondInsert != 2 {
		t.Errorf("RegisterPost tests fail")
	}
	os.Remove("./test.db")
}
