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
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	var post config.NewPost
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	res.Header().Set("Content-Type", "application/json")

	err := decoder.Decode(&post)
	if err != nil {
		log.Println("Error decoding post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload",
		}
		json.NewEncoder(res).Encode(failed)
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterPost(db, post.Title, post.Datas, post.IdUser)
	if err != nil {
		log.Println("Error register post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while saving post",
		}
		json.NewEncoder(res).Encode(failed)
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The post has been saved on id: " + fmt.Sprint(id),
		Id:      id,
	}
	json.NewEncoder(res).Encode(successfull)
}
