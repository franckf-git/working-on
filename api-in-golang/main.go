package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WelcomePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "welcome")
}

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "all posts")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", WelcomePage)
	router.HandleFunc("/api/v1/posts", ShowAllPosts)
	log.Fatal(http.ListenAndServe(":8000", router))
}
