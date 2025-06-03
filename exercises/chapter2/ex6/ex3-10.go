package main

import (
	"bytes"
	"fmt"
)

func reverse(b []byte) []byte {
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return b
}

// 1
// 1,234,895
// 1234895
func newComma(s string) string {
	var buf bytes.Buffer

	n := len(s)
	count := 0

	for i := n - 1; i >= 0; i-- {
		if count != 0 && count%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])
		count++
	}
	return string(reverse(buf.Bytes()))
	
}

func main() {
	fmt.Println(newComma("1234895"))
}
