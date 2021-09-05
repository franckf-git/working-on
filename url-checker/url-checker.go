package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"http://localhost:5051/posts",
	"http://localhost:5052/posts",
	"https://jsonplaceholder.typicode.com/todos/",
	"https://www.twitch.tv/",
}

type urlsStatus struct {
	mu      sync.Mutex
	results map[string]bool
}

func (exec *urlsStatus) checkUrl(url string) {
	exec.mu.Lock()
	resp, err := http.Get(url)
	exec.results[url] = true
	if err != nil {
		fmt.Println("error accessing url:", url, err)
		exec.results[url] = false
	}
	if resp.StatusCode > 299 {
		fmt.Println("Response failed with status code:", resp.StatusCode)
		exec.results[url] = false
	}
	exec.mu.Unlock()
}

func main() {
	checker := urlsStatus{results: make(map[string]bool)}
	for _, url := range urls {
		checker.checkUrl(url)
	}
	fmt.Println(checker.results)
}
