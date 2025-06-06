package main

import (
	"fmt"
	"strings"
)

func main() {
	split := "-"

	fmt.Println(joinNew(split, "ivan", "vivaldi", "azis"))
}

func joinNew(separator string, names ...string) string {
	var sb strings.Builder

	for i, v := range names {
		sb.WriteString(v)
		
		if i == len(names) - 1 {
			break
		}
		
		sb.WriteString(separator)
	}

	return sb.String()
}


