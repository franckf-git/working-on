package controllers

import (
	"fmt"
	"net/http"
)

func ShowAllPosts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "all posts")
}
