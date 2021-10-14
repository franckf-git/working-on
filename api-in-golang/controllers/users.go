package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil {
		config.ErrorLogg("AddUser(controllers) - decoding user:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(res)
		return
	}

	if !CheckEmailPassword(user) {
		config.ErrorLogg("AddUser(controllers) - email/password valisator:", user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "error in email or password validator - email must be a valid email and password must be at least 8 characters, uppercase, lowercase, numbers and specials included",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusPreconditionRequired)
		w.Write(res)
		return
	}

	id, err := user.RegisterUser(Db)
	if err != nil {
		config.ErrorLogg("AddUser(controllers) - register user:", err, user.Email)
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

func AskJWT(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		config.ErrorLogg("AskJWT(controllers) - decoding user:", err)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(res)
		return
	}

	id, err := user.CheckExistingUser(Db)
	if err != nil {
		config.ErrorLogg("AskJWT(controllers) - logging user:", err, user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "This email doesn't exist or the password is wrong",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}

	tokenString, err := GenerateToken(id)
	if err != nil {
		config.ErrorLogg("AskJWT(controllers) - signin token:", err, tokenString)
		failed := config.Message{
			Status:  "error",
			Message: "error while signin token",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
		return
	}

	successfull := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		JWT     string `json:"jwt"`
	}{
		Status:  "success",
		Message: "Successfull auth, JWT created, it is valid for 24H",
		JWT:     tokenString,
	}
	res, _ := json.Marshal(successfull)
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
