package main

import (
	"fmt"
	// "time"
)

func main() {
	go fmt.Println("Hello from goroutine 1")
	go fmt.Println("Hello from goroutine 2")
	fmt.Println("Hello from main")
	// time.Sleep(1 * time.Second)
}
