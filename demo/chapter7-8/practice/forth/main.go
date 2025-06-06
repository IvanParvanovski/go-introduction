package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		go func(n int) {
			defer wg.Done()
			fmt.Println("Running task", n)
		}(i)
	}

	wg.Wait()
}