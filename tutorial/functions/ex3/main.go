package main 

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Stacey"
	
	sayGreeting(&greeting, &name)

	fmt.Println(name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	
	*name = "Ted"
	
	fmt.Println(*name)
}

