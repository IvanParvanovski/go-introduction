package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}
}
func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

