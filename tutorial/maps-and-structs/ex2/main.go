package main

import (
	"fmt"
)

type Doctor struct {
	number int
	actorName string
	companions []string
	episodes []string
}

// this wold export the doctor struct and 
// you will be able to access the field names
// type Doctor struct {
// 	Number int
// 	ActorName string
// 	Companions []string
// 	Episodes []string
// }



func main() {
	aDoctor := Doctor {
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},

	}	
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[1])
}
