package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Println("Error bad content-type formating:", r.Header)
		failed := config.Message{
			Status:  "error",
			Message: "error bad content-type formating:" + fmt.Sprint(r.Header),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write(res)
		return
	}

	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error decoding user:", err, user)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err, user),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(res)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := user.RegisterUser(db)
	if err != nil {
		log.Println("Error register user:", err, user)
		failed := config.Message{
			Status:  "error",
			Message: "error while saving user",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	successfull := config.Message{
		Status:  "success",
		Message: "The user has been saved on id: " + fmt.Sprint(id),
		Id:      id,
	}
	res, _ := json.Marshal(successfull)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
