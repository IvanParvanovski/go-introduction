package main 

import (
	"fmt"
)

func main() {
	// this is array
	a := [...]int{1, 2, 3}
	// array passed by reference
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}