package main 

import (
	"fmt"
)

var (
	actorName string = "Elisabeth Sladen"
	companion string = "Sarah Jane Smith"
	doctorName int = 3
	season int = 11
)

var (
	counter int = 0
)

var i int = 27

func main() {
	fmt.Println(i)

	var i int = 42
	// i = 13 <- not possible, already declared
	
	// Shadowing
	fmt.Println(i)
}
