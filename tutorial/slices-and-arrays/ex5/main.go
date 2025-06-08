package main 

import (
	"fmt"
)

func main() {
	// this is a slice
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	// if this is an array it would be a copy
	// if this is a slice it is passed as a pointer
	
}
