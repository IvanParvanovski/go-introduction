package main 

import (
	"fmt"
)

func main() {
	// go copies the entire array
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5 
	fmt.Println(a)
	fmt.Println(b)
}
