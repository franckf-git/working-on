package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
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
	if err != nil || user.Email == "" || user.Password == "" {
		log.Println("Error decoding user:", err, user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err, user.Email),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(res)
		return
	}

	if !CheckEmailPassword(user) {
		log.Println("Error in email or password validator:", user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "error in email or password validator - email must be a valid email and password must be at least 8 characters, uppercase, lowercase, numbers and specials included",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusPreconditionRequired)
		w.Write(res)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := user.RegisterUser(db)
	if err != nil {
		log.Println("Error register user:", err, user.Email)
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

func CheckEmailPassword(user models.User) bool {
	var validEmail = regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)
	if !validEmail.MatchString(user.Email) {
		return false
	}

	if len(user.Password) < 8 {
		return false
	}

	var numbers = regexp.MustCompile("[0-9]")
	var lower = regexp.MustCompile("[a-z]")
	var upper = regexp.MustCompile("[A-Z]")
	var nospecials = regexp.MustCompile(`[^\w]`)

	if !numbers.MatchString(user.Password) {
		return false
	}
	if !lower.MatchString(user.Password) {
		return false
	}
	if !upper.MatchString(user.Password) {
		return false
	}
	if !nospecials.MatchString(user.Password) {
		return false
	}
	return true
}

func AskJWT(w http.ResponseWriter, r *http.Request) {
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
	if err != nil || user.Email == "" || user.Password == "" {
		log.Println("Error decoding user:", err, user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "error while decoding payload " + fmt.Sprint(err, user.Email),
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write(res)
		return
	}

	db := models.OpenDatabase()
	defer db.Close()
	id, err := user.CheckExistingUser(db)
	if err != nil {
		log.Println("Error logging user:", err, user.Email)
		failed := config.Message{
			Status:  "error",
			Message: "This email doesn't exist or the password is wrong",
		}
		res, _ := json.Marshal(failed)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(res)
		return
	}

	var hmacKey = []byte(config.JWTkey)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"userId": id,
		"expire": expiresAt,
	})
	tokenString, err := token.SignedString(hmacKey)
	if err != nil {
		log.Println("Error signin token:", err, tokenString)
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
