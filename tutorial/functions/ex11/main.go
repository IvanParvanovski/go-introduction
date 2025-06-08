package main
import "fmt"

type counter int

func main() {
	c := counter(0)

	fmt.Println(c)	
	c.increaseCount()
	fmt.Println(c)
}

func (c *counter) increaseCount() {
	for i := 0; i < 10; i++ {
		*c++
	}
}
