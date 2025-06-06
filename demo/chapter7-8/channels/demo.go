package main
import "fmt"

func main() {
	number := make(chan int)
	message := make(chan string)

	go channelData(number, message)

	fmt.Println("Channel Data:", <-number)
	fmt.Println("Channel Data:", <-number)
	fmt.Println("Channel Data:", <-message)
}

func channelData(number chan int, message chan string) {
	number <- 15
	number <- 16
	message <- "Learning Go channel"
}
