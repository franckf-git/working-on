package router

import (
	"encoding/json"
	"fmt"
	"lite-api-crud/config"
	"lite-api-crud/controllers"
	"net/http"
)

func isAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
		_, errJWT := controllers.ValidateToken(authToken)
		// work but the jwt must be decrypt twice - one for auth and one to get idUser
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
		next.ServeHTTP(w, r)
	}
}

func setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
