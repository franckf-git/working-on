package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/models"
	"net/http"
)

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	posts := models.GetAllPosts(db)
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterPost(db, "title", "datas", 1) // test entry
	fmt.Fprintln(res, "add post", id, err)
}
