package main 

import (
	"fmt"
	"time"
)

func main() {
	// There is a dependency between 
	// the main stack and the goroutine
	// this is race condition
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}