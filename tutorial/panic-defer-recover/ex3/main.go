package main 

import (
	"fmt"
)

func main() {
	// it saves the state
	a := "start"
	defer fmt.Println(a)
	a = "end"
}
