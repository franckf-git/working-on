package tests

import (
	"database/sql"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"os"
	"reflect"
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

func Test_GetAllPosts(t *testing.T) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal("Fail to open database:", err)
	}
	defer db.Close()
	models.StartDatabase(db)
	models.RegisterPost(db, "title1", "datas1", 1)
	models.RegisterPost(db, "title2", "datas2", 2)
	postsTests := models.GetAllPosts(db)
	postsTests[0].Created = ""
	postsTests[1].Created = ""
	want := []config.Post{
		{1, "title1", "datas1", "", 1},
		{2, "title2", "datas2", "", 2},
	}
	if !reflect.DeepEqual(postsTests, want) {
		t.Errorf("GetAllPosts tests fail")
	}
	os.Remove("./test.db")
}
