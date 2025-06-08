package main

import (
	"fmt"
)

func main() {
	// remove an element from the beginning and at the end
	a := []int{1, 2, 3, 4, 5}
	b := a[1:len(a) - 1]
	fmt.Println(b)

	// remove an element from the middle
	c := append(a[:2], a[3:]...)
	fmt.Println(c)
}

