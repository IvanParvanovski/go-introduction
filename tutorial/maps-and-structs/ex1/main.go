package main 

import (
	"fmt"
)

func main() {
	statePopulations := make(map[string]int)
	
	statePopulations = map[string]int {
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11614373,
	}
	// m := map[[3]int]string{}
	// fmt.Println(statePopulations, m)

	// delete(statePopulations, "Georgia")
	// pop, ok := statePopulations["Ohio"]
	// fmt.Println(pop, ok)
	
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}
