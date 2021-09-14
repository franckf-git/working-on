package controllers

import (
	"fmt"
	"lite-api-crud/models"
	"net/http"
)

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "all posts")
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterPost(db, "title", "datas", 1) // test entry
	fmt.Fprintln(res, "add post", id, err)
}
