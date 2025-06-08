package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
	ch := make(chan int)
	
	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func(ch <-chan int) {
			i := <-ch 
			fmt.Println(i)
			// ch <- 27 <-- you cannot save data because
			// of the function definition 
			wg.Done()
		}(ch)
		
		go func(ch chan<- int) {
			ch <- 42 // <-- you can only write data to the
			// functoin because of the definition
			wg.Done()
		}(ch)
	}

	wg.Wait()
}
