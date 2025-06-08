package main 

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
func main() {
	ch := make(chan int)

	// go func() {
	// 	i := <-ch 
	// 	fmt.Println(i)
	// 	wg.Done()
	// }()
	
	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch 
			fmt.Println(i)
			ch <- 27
			wg.Done()
		}()

		go func() {
			ch <- 42
			fmt.Println(<-ch)
			wg.Done()
		}()
	}

	wg.Wait()
}
