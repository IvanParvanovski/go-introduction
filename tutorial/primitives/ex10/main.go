package main 

import (
	"fmt"
)

func main() {
	s := "this is a string"
	b := []byte(s)

	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	fmt.Printf("%v, %T\n", b, b)
}

