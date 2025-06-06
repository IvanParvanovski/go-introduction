package main

import (
	"fmt"
	"sync"
)

var balance int
var mu sync.Mutex

type Bank struct {
	money int
}

func (b Bank) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	balance = balance + amount
	fmt.Println("=", balance) 
	mu.Unlock()
}

func (b Bank) Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	bank := Bank{500}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		bank.Deposit(200, &wg)
	}()

	go bank.Deposit(100, &wg)

	wg.Wait()
}
