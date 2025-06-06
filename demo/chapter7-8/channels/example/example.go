package main 

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string, ch chan<- string){
	start := time.Now()
	resp, err := http.Get(url)
	
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	elapsed := time.Since(start)
	
	ch <- fmt.Sprintf("Fetched %s in %v (Status: %s)", url, elapsed, resp.Status)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://httpbin.org/delay/2",
	}

	ch  := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Print(<-ch)
	}
}