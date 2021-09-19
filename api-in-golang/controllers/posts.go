package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		res.WriteHeader(http.StatusUnsupportedMediaType)
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
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(failed)
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The post has been saved on id: " + fmt.Sprint(id),
		Id:      id,
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(successfull)
}

func ShowPost(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	post := models.GetPost(db, id)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(post)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	var post config.NewPost
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	res.Header().Set("Content-Type", "application/json")

	err := decoder.Decode(&post)
	if err != nil {
		log.Println("Error decoding post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload",
		}
		res.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(res).Encode(failed)
	}

	db := models.OpenDatabase()
	defer db.Close()
	err = models.UpdatingPost(db, id, post.Title, post.Datas, post.IdUser)
	if err != nil {
		log.Println("Error updating post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while updating post",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(failed)
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The post has been updated on id: " + fmt.Sprint(id),
		Id:      id,
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(successfull)
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	db := models.OpenDatabase()
	defer db.Close()
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	res.Header().Set("Content-Type", "application/json")
	err := models.DeletingPost(db, id)
	if err != nil {
		log.Println("Error deleting post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while deleting post",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(failed)
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The post has been deleted",
		Id:      id,
	}
	json.NewEncoder(res).Encode(successfull)
}
