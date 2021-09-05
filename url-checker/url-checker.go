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

func (exec *urlsStatus) checkUrl(url string, urlChannel chan map[string]bool) {
	result := make(map[string]bool)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error accessing url:", url, err)
		result[url] = false
		urlChannel <- result
	} else {
		if resp.StatusCode > 299 {
			fmt.Println("Response failed with status code:", resp.StatusCode)
			result[url] = false
			urlChannel <- result
		} else {
			result[url] = true
			urlChannel <- result
		}
	}
}

func main() {
	checker := urlsStatus{results: make(map[string]bool)}
	urlChannel := make(chan map[string]bool)
	for _, url := range urls {
		checker.mu.Lock()
		go checker.checkUrl(url, urlChannel)
		checker.mu.Unlock()
	}
	// double loop: one to get resutl form chan and one to fusion two maps
	for i := 0; i < len(urls); i++ {
		result := <-urlChannel
		for k, v := range result {
			checker.results[k] = v
		}
	}
	fmt.Println(checker.results)
}

/*
// With waitgroup - not async because of mutex
func main() {
	checker := urlsStatus{results: make(map[string]bool)}
	var waitForAllUrls sync.WaitGroup
	for _, url := range urls {
		waitForAllUrls.Add(1)
		go func(asyncUrl string) {
			defer waitForAllUrls.Done()
			checker.checkUrl(asyncUrl)
			}(url)
		}
		waitForAllUrls.Wait()
		fmt.Println(checker.results)
	}
*/
