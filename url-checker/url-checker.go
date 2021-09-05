package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://localhost:5051/posts",
	"http://localhost:5052/posts",
	"https://jsonplaceholder.typicode.com/todos/",
	"https://www.twitch.tv/",
}

func main() {
	var results = make(map[string]bool)
	results["test"] = true
	checkUrl(urls[0])
	fmt.Println(results)
}

func checkUrl(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error accessing url:", url, err)
		return false
	}
	if resp.StatusCode > 299 {
		fmt.Println("Response failed with status code:", resp.StatusCode)
		return false
	}
	return true
}
