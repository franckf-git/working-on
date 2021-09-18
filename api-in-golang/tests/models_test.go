package tests

import (
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	models.InitializeDB()
}

var fakeCreatedTime string = time.Now().Format(time.RFC3339)

func Test_RegisterPost(t *testing.T) {
	db := models.OpenDatabase()
	defer db.Close()
	models.CleanTables(db)

	firstInsert, _ := models.RegisterPost(db, "title1", "datas1", 1)
	secondInsert, _ := models.RegisterPost(db, "title2", "datas2", 2)
	if firstInsert != 1 && secondInsert != 2 {
		t.Errorf("RegisterPost tests fail")
	}
}

func Test_GetAllPosts(t *testing.T) {
	db := models.OpenDatabase()
	defer db.Close()

	postsTests := models.GetAllPosts(db)
	postsTests[0].Created = fakeCreatedTime
	postsTests[1].Created = fakeCreatedTime
	want := []config.Post{
		{Id: 1, Title: "title1", Datas: "datas1", Created: fakeCreatedTime, IdUser: 1},
		{Id: 2, Title: "title2", Datas: "datas2", Created: fakeCreatedTime, IdUser: 2},
	}
	if !reflect.DeepEqual(postsTests, want) {
		fmt.Println(postsTests)
		t.Errorf("GetAllPosts tests fail")
	}
}
