package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	a2 = iota
)

func main() {
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	
	fmt.Printf("%v\n", a2)
}
