package main 

import (
	"fmt"
)

func main() {
	var a int = 42 
	var b *int = &a 
	fmt.Println(&a, b)
	fmt.Println(a, *b)

	a =  27 

	fmt.Println(a, *b)
}
