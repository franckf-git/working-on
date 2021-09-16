package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"net/http"
)

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	posts := models.GetAllPosts(db)
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	var post config.NewPost
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()

	err := decoder.Decode(&post)
	if err != nil {
		log.Println("Error decoding post:", err, post)
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterPost(db, post.Title, post.Datas, post.IdUser)
	fmt.Fprintln(res, "add post", id, err)
}
