package tests

import (
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"lite-api-crud/models"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var fakeCreatedTime string = time.Now().Format(time.RFC3339)

func Test_RegisterPost(t *testing.T) {
	firstInsert, _ := models.RegisterPost(controllers.Db, "title1", "datas1", 1)
	secondInsert, _ := models.RegisterPost(controllers.Db, "title2", "datas2", 2)
	if firstInsert != 1 && secondInsert != 2 {
		t.Errorf("RegisterPost tests fail")
	}
}

func Test_GetAllPosts(t *testing.T) {
	postsTests, _ := models.GetAllPosts(controllers.Db)
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
