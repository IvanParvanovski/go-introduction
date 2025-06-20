package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// ch := make(chan int, 50)
	ch := make(chan int)
	
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}

		wg.Done() 
	}(ch)
	go func(ch chan<- int) {
		ch <- 42 
		ch <- 27
		close(ch)
		
		wg.Done()
	}(ch)
	wg.Wait()
}