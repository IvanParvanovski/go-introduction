package main 

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	var j float32 = 27
	k := 99

	fmt.Printf("%v, %T\n", j, j)
	fmt.Println(i)
	fmt.Println(k)
}
