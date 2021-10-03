package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
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

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		config.ErrorLogg("AddPost(controllers) - bad content-type formating:", req.Header)
		failed := config.Message{
			Status:  "error",
			Message: "error bad content-type formating:" + fmt.Sprint(req.Header),
		}
		res.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(res).Encode(failed)
		return
	}

	authToken := req.Header.Get("Authorization")
	idUserJWT, errJWT := validateToken(authToken)
	if errJWT != nil {
		config.ErrorLogg("AddPost(controllers) - decoding JWT:", errJWT)
		failed := config.Message{
			Status:  "error",
			Message: "error decoding JWT:" + fmt.Sprint(errJWT),
		}
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(failed)
		return
	}

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&post)
	if err != nil {
		config.ErrorLogg("AddPost(controllers) - decoding post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err, post),
		}
		res.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(res).Encode(failed)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterPost(db, post.Title, post.Datas, idUserJWT)
	if err != nil {
		config.ErrorLogg("AddPost(controllers) - register post:", err, post)
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
	post, err := models.GetPost(db, id)
	if err != nil {
		config.ErrorLogg("ShowPost(controllers) - getting post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while getting post " + fmt.Sprint(err),
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(failed)
		return
	}
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

	contentType := req.Header.Get("Content-Type")
	if contentType != "application/json" {
		config.ErrorLogg("UpdatePost(controllers) - bad content-type formating:", req.Header)
		failed := config.Message{
			Status:  "error",
			Message: "error bad content-type formating:" + fmt.Sprint(req.Header),
		}
		res.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(res).Encode(failed)
		return
	}

	authToken := req.Header.Get("Authorization")
	idUserJWT, errJWT := validateToken(authToken)
	if errJWT != nil {
		config.ErrorLogg("UpdatePost(controllers) - decoding JWT:", errJWT)
		failed := config.Message{
			Status:  "error",
			Message: "error decoding JWT:" + fmt.Sprint(errJWT),
		}
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(failed)
		return
	}

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&post)
	if err != nil {
		config.ErrorLogg("UpdatePost(controllers) - decoding post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload" + fmt.Sprint(err, post),
		}
		res.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(res).Encode(failed)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	ids, err := models.GetAllPostsByUser(db, idUserJWT)
	if !find(ids, id) {
		config.ErrorLogg("UpdatePost(controllers) - this user can't update this post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error this user can't update this post",
		}
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(failed)
		return
	}

	err = models.UpdatingPost(db, id, post.Title, post.Datas, idUserJWT)
	if err != nil {
		config.ErrorLogg("UpdatePost(controllers) - updating post:", err, post)
		failed := config.Message{
			Status:  "error",
			Message: "error while updating post " + fmt.Sprint(err),
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(failed)
		return
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
	authToken := req.Header.Get("Authorization")
	idUserJWT, errJWT := validateToken(authToken)
	if errJWT != nil {
		config.ErrorLogg("DeletePost(controllers) - decoding JWT:", errJWT)
		failed := config.Message{
			Status:  "error",
			Message: "error decoding JWT:" + fmt.Sprint(errJWT),
		}
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(failed)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	res.Header().Set("Content-Type", "application/json")

	ids, err := models.GetAllPostsByUser(db, idUserJWT)
	if !find(ids, id) {
		config.ErrorLogg("DeletePost(controllers) - this user can't delete this post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error this user can't delete this post",
		}
		res.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(res).Encode(failed)
		return
	}

	err = models.DeletingPost(db, id)
	if err != nil {
		config.ErrorLogg("DeletePost(controllers) - deleting post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while deleting post",
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(failed)
		return
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The post has been deleted",
		Id:      id,
	}
	json.NewEncoder(res).Encode(successfull)
}
