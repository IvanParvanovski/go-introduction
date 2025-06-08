package main

import (
	"fmt"
)

func main() {
	// const a int = 42
	const a = 42
	const b int = 27

	var c int16 = 27
	fmt.Printf("%v, %T\n", a + b, a + b)
	fmt.Printf("%v, %T\n", a + c, a + c)
}