package main

import (
	"fmt"
	"sync"
)


func main() {
	
}

func mutexSolution() {
	m := make(map[string]int)
	
	mu := sync.Mutex{}
	go func() {
		mu.Lock()
		m["foo"] = 1
		mu.Unlock()
	}()


	mu.Lock()
	m["bar"] = 2 
	for k, v := range m {
		fmt.Println(k, v)
	}
	mu.Unlock()

}

func channelSolution() {
	m := make(map[string]int)
	
	ch := make(chan bool)

	go func() {
		m["foo"] = 1

		ch <- true
	}()

	<- ch
	m["bar"] = 2 
	for k, v := range m {
		fmt.Println(k, v)
	}
}