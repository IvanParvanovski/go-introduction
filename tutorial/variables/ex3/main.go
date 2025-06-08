package main

import (
	"fmt"
)

func main() {
	var i int = 69
	fmt.Printf("%v, %T\n", i, i)

	var j string 
	j = string(i) // <- convert from unicode number 69
	fmt.Printf("%v, %T\n", j, j)
}
