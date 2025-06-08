package main 

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	
	return n, err
}


