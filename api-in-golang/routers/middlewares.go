package router

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"net/http"
)

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Authorization"] == nil {
			config.ErrorLogg("isAuthorized(routes) - no authorization token found")
			failed := config.Message{
				Status:  "error",
				Message: "no authorization token found",
			}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(failed)
			return
		}

		authToken := r.Header.Get("Authorization")
		idUser, errJWT := controllers.ValidateToken(authToken)
		if errJWT != nil {
			config.ErrorLogg("isAuthorized(routes) - decoding JWT:", errJWT)
			failed := config.Message{
				Status:  "error",
				Message: "error decoding JWT:" + fmt.Sprint(errJWT),
			}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(failed)
			return
		}
		r.Header.Set("idUser", fmt.Sprint(idUser))
		next.ServeHTTP(w, r)
	})
}

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func checkContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			config.ErrorLogg("AddPost(controllers) - bad content-type formating:", r.Header)
			failed := config.Message{
				Status:  "error",
				Message: "error bad content-type formating:" + fmt.Sprint(r.Header),
			}
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(failed)
			return
		}
		next.ServeHTTP(w, r)
	})

}
