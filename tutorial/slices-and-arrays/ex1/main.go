package main 

import (
	"fmt"
)

func main() {
	grades := [...]int{97, 85, 93}
	// grades := [3]int{97, 85, 93}

	fmt.Printf("Grades: %v\n", grades)

	var students [3]string 
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[2] = "Michael"
	students[1] = "Ahmed"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of Students: %v\n", len(students))
}

