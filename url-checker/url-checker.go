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

type urlOK struct {
	url    string
	status bool
}

func checkUrl(url string, urlChannel chan urlOK) {
	result := urlOK{}
	result.url = url
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error accessing url:", url, err)
		result.status = false
		urlChannel <- result
	} else {
		if resp.StatusCode > 299 {
			fmt.Println("response failed with status code:", resp.StatusCode)
			result.status = false
			urlChannel <- result
		} else {
			result.status = true
			urlChannel <- result
		}
	}
}

func main() {
	urlChannel := make(chan urlOK)
	for _, url := range urls {
		go checkUrl(url, urlChannel)
	}

	results := make([]urlOK, 0)
	for i := 0; i < len(urls); i++ {
		result := <-urlChannel
		results = append(results, result)
	}
	fmt.Println(results)
}
