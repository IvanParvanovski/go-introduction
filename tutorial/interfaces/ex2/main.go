package main 

import (
	"fmt"
)

type IntCounter int 

type Incrementer interface {
	Increment() int
}

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
