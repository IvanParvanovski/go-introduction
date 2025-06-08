package main 

import (
	"fmt"
)

func main() {
	switch i := 2 + 3; i {
		case 1, 3, 5: 
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("not one or two")
	}
}