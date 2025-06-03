// go run ex3/ex1-3.go asd asd

package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Slow implementation:")
	measureTime(slowImplementation)
	
	fmt.Println()
	
	fmt.Println("Fast implementation:")
	measureTime(fastImplementation)
}

func measureTime(f func()) { 
	start := time.Now()
	f()
	end := time.Now()
	
	diff := end.Sub(start)
	fmt.Println(fmt.Sprintf("Time it took: %s", diff))
}

func slowImplementation() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func fastImplementation() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
