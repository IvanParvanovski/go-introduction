package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello from gorouting"
	}()

	msg := <-ch
	fmt.Println(msg)
}

