package controllers

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func WelcomePage(res http.ResponseWriter, req *http.Request) {
	homepage := config.Message{
		Status:  "information",
		Message: config.WelcomeMessage,
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(homepage)
}

func Docs(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, config.DocsLink, http.StatusMovedPermanently)
}

func NotFoundMessage(res http.ResponseWriter, req *http.Request) {
	notfound := config.Message{
		Status:  "error",
		Message: "this route doesn't exist",
	}
	res.WriteHeader(http.StatusNotFound)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(notfound)
}

func validateToken(authToken string) (idUser int, err error) {
	parseForAuthToken := strings.Split(authToken, " ")
	if len(parseForAuthToken) != 2 {
		return 0, fmt.Errorf("bad formating in Bearer")
	}
	tokenToValidate := parseForAuthToken[1]

	var claims config.JwtInfos
	token, err := jwt.ParseWithClaims(tokenToValidate, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWTkey), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("token is invalid")
	}
	idUser = claims.IdUser
	expiresAt := claims.ExpiresAt
	if expiresAt < time.Now().UTC().Unix() {
		return 0, fmt.Errorf("token is expire")
	}
	return
}
