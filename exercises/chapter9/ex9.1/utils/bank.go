package bank

// Package bank provides a concurrency-safe bank with one account.
var deposits = make(chan int) 
var balances = make(chan int) 
var withdrawals = make(chan withdrawRequest)

type withdrawRequest struct {
	amount int 
	success chan bool
}

func Deposit(amount int) { 
	deposits <- amount 
}


func Balance() int { 
	return <-balances 
}
func Withdraw(amount int) bool {
	result := make(chan bool)
	withdrawals <- withdrawRequest{amount, result}
	return <-result
}


func teller() {
	var balance int 

	for {
		select {
			case amount := <-deposits:
				balance += amount
			case balances <- balance:
			
			case req := <-withdrawals:
				if req.amount <= balance {
					balance -= req.amount
					req.success <- true
				} else {
					req.success <- false
				}
		}
	}
}
func init() {
	go teller()
}

