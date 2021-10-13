package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	posts, err := models.GetAllPosts(Db)
	if err != nil {
		config.ErrorLogg("ShowAllPostsPost(controllers) - getting posts:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while retriving posts",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(failed)
		return
	}
	if len(posts) == 0 {
		failed := config.Message{
			Status:  "success",
			Message: "no posts in database",
		}
		json.NewEncoder(res).Encode(failed)
		return
	}
	json.NewEncoder(res).Encode(posts)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	post := decodePayloadPost(res, req.Body)
	idUserJWT, _ := strconv.Atoi(req.Header.Get("idUser"))

	id, err := models.RegisterPost(Db, post.Title, post.Datas, idUserJWT)
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
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	post, err := models.GetPost(Db, id)
	if err != nil {
		config.ErrorLogg("ShowPost(controllers) - getting post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while getting post " + fmt.Sprint(err),
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(failed)
		return
	}
	json.NewEncoder(res).Encode(post)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	post := decodePayloadPost(res, req.Body)
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	idUserJWT, _ := strconv.Atoi(req.Header.Get("idUser"))

	ids, err := models.GetAllPostsByUser(Db, idUserJWT)
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

	err = models.UpdatingPost(Db, id, post.Title, post.Datas, idUserJWT)
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
	idUserJWT, _ := strconv.Atoi(req.Header.Get("idUser"))

	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	ids, err := models.GetAllPostsByUser(Db, idUserJWT)
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

	err = models.DeletingPost(Db, id)
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

func decodePayloadPost(w http.ResponseWriter, body io.ReadCloser) (post config.Post) {
	decoder := json.NewDecoder(body)
	defer body.Close()
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&post)
	if err != nil {
		config.ErrorLogg("(controllers) - decoding post:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err),
		}
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(failed)
		return
	}
	return
}
