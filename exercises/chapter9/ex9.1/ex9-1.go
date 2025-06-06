package main

import (
	"ex9-1/utils"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		bank.Deposit(200)
		fmt.Println("Deposited 200")
		wg.Done()
	}()

	go func() {
		ok := bank.Withdraw(100)
		fmt.Println("Withdraw 100 success:", ok)
		wg.Done()
	}()

	go func() {
		ok := bank.Withdraw(150)
		fmt.Println("Withdraw 150 success:", ok)
		wg.Done()
	}()

	go func() {
		balance := bank.Balance()
		fmt.Println("Final balance:", balance)
		wg.Done()
	}()

	wg.Wait()
}


