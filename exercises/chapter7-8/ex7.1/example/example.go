package main 

import (
	"fmt"
)

type ByteCounter int 

type Stringer interface {
	String()string 
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	fmt.Println(p)
	
	*c += ByteCounter(len(p))
	
	return len(p), nil
}


func main() {
	var c ByteCounter
	c.Write([]byte("hållo"))

	fmt.Println(c)
	c = 0
	
	var name = "Dolly"
	
	fmt.Fprintf(&c, "hello %s", name)
	fmt.Println(c)
}
