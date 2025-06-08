package main

import (
	"fmt"
)

func main() {
	n := 42
	var m uint16 = 42 // <-- uint are only positive
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)
}
