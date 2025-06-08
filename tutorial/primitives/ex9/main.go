package main 

import (
	"fmt"
)

func main() {
	var n complex128 = 1 + 2i
	fmt.Printf("%v, %T\n", real(n), real(n))
	fmt.Printf("%v, %T\n", n, n)
}