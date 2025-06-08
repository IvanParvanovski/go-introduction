package main 

import (
	"fmt"
)

func main() {
	sayGreeting("Hello", "Stacey")
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	
} 
