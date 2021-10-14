package tests

import (
	"database/sql"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var fakeCreatedTime string = time.Now().Format(time.RFC3339)

var db *sql.DB = models.InitializeDB("test")

func Test_RegisterPost(t *testing.T) {
	firstInsert, _ := models.RegisterPost(db, "title1", "datas1", 1)
	secondInsert, _ := models.RegisterPost(db, "title2", "datas2", 2)
	if firstInsert != 1 && secondInsert != 2 {
		t.Errorf("RegisterPost tests fail")
	}
}

func Test_GetAllPosts(t *testing.T) {
	postsTests, _ := models.GetAllPosts(db)
	postsTests[0].Created = fakeCreatedTime
	postsTests[1].Created = fakeCreatedTime
	want := []config.GetPost{
		{
			Id: 1,
			Post: config.Post{
				Title:  "title1",
				Datas:  "datas1",
				IdUser: 1,
			},
			Created: fakeCreatedTime,
		},
		{
			Id: 2,
			Post: config.Post{
				Title:  "title2",
				Datas:  "datas2",
				IdUser: 2,
			},
			Created: fakeCreatedTime,
		},
	}
	if !reflect.DeepEqual(postsTests, want) {
		fmt.Println(postsTests)
		t.Errorf("GetAllPosts tests fail")
	}
}
