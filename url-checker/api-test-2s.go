package main

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"time"
)

//go:embed spacex-example.json
var posts string

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	json.NewEncoder(w).Encode(posts)
}

func main() {
	http.HandleFunc("/posts", PostsHandler)
	http.ListenAndServe(":5052", nil)
}
