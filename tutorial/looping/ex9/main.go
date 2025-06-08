package main 

import (
	"fmt"
)

func main() {
	s := "hello Go!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
	
	fmt.Println()

	statePopulations := map[string]int {
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11514373,
	}


	// get only keys
	for k := range statePopulations {
		fmt.Println(k)
	}

	fmt.Println()

	// get only values
	for _, v := range statePopulations {
		fmt.Println(v)
	}
}
