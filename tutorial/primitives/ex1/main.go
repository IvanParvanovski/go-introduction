package main

import (
	"fmt"
)

func main() {
	// var n bool = true
	n := 1 == 1
	m := 1 == 2
	var z bool

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", z, z)
}

