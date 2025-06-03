package main 

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)

	
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		if counts[filename] == nil {
			counts[filename] = make(map[string]int)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[filename][line]++
		}
	}


	for k1, v1 := range counts {
		fmt.Printf("%s\n", k1)

		for k2, v2 := range v1 {
			if (v2 > 1) {
				fmt.Printf("\t%s -> %d\n", k2, v2)
			}
		}
	}

	// for line, n := range counts {
	// 	fmt.Printf("%d\t%s\n", n, line)
	// 	// if n > 1 {
	// 	// 	fmt.Printf("%d\t%s\n", n , line)
	// 	// }
	// }
}