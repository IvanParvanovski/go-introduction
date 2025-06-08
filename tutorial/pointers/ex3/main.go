package main 

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main() {
	var ms *myStruct 
	fmt.Println(ms)

	ms = &myStruct{foo: 42}
	fmt.Println(ms)

	ms.foo = 42
	fmt.Println(ms.foo)
}
