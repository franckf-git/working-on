package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	/*
		body, _ := ioutil.ReadAll(r.Body)
		var user models.User
		json.Unmarshal(body, &user)
		log.Printf("user: %#+v\n", user)
	*/

	user := models.User{
		Email:    "dfsg",
		Password: "sdfg",
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := models.RegisterUser(db, user)
	if err != nil {
		log.Printf("err: %#+v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	successfull := config.Message{
		Status:  "success",
		Message: "The user has been saved on id: " + fmt.Sprint(id),
		Id:      id,
	}
	res, _ := json.Marshal(successfull)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
	return
}
