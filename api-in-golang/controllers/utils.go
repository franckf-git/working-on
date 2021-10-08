package controllers

import (
	"database/sql"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/models"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var Db *sql.DB = models.InitializeDB(config.State)

func ValidateToken(authToken string) (idUser int, err error) {
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

func GenerateToken(id int) (string, error) {
	var hmacKey = []byte(config.JWTkey)
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	claims := config.JwtInfos{
		IdUser:         id,
		ExpiresAt:      expiresAt,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(hmacKey)
	return tokenString, err
}

func CheckEmailPassword(user models.User) bool {
	user.Email = strings.Trim(user.Email, " ")
	var validEmail = regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)
	if !validEmail.MatchString(user.Email) {
		return false
	}

	var newline = regexp.MustCompile("\n")
	if newline.MatchString(user.Email) {
		return false
	}

	if len(user.Password) < 8 {
		return false
	}

	if len(user.Password) > 52 || len(user.Email) > 52 {
		return false
	}

	var numbers = regexp.MustCompile("[0-9]")
	var lower = regexp.MustCompile("[a-z]")
	var upper = regexp.MustCompile("[A-Z]")
	var nospecials = regexp.MustCompile(`[^\w]`)
	var nowhitepace = regexp.MustCompile(" ")

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
	if nowhitepace.MatchString(user.Password) {
		return false
	}
	return true
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
