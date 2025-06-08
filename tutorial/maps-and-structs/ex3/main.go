package main 

import (
	"fmt"
)

type Doctor struct {
	name string
}

func main() {
	// anonymous struct
	aDoctor := struct{name string}{name: "John Pertwee"}
	
	// THIS IS A COPY !!!
	// anotherDoctor := &aDoctor
	anotherDoctor := aDoctor
	anotherDoctor.name = "Tom Baker"

	fmt.Println(anotherDoctor)
	fmt.Println(aDoctor)

}
