package main

import (
	"fmt"
)

func main() {
	a := []int{}
	
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 1)
	
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	a = append(a, 2, 3, 4, 5)

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	// very useful!!!
	a = append(a, []int{ 2, 3, 4, 5}...)

	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

}
