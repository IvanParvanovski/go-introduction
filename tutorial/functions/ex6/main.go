package main

import (
	"fmt"
)

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", s)
}

func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}