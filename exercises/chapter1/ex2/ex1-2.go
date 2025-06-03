package main 

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(fmt.Sprintf("index %d: %s", i, os.Args[i]))
	}
}