package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	fmt.Printf(strings.Join(os.Args, " "))
}
