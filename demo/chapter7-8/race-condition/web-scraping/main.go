package main 

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make(map[string]int)

	urls := []string {
		"https://example.com",
		"https://golang.org",
		"https://httpbin.org/status/404",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg, &mu, results)
	}

	wg.Wait()

	for url, status := range results {
		fmt.Printf("%s -> %d\n", url, status)
	}
}

func fetch(url string, wg *sync.WaitGroup, mu *sync.Mutex, results map[string]int ) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching:", url)
		return
	}
	defer resp.Body.Close()

	mu.Lock()
	results[url] = resp.StatusCode
	mu.Unlock()
}
