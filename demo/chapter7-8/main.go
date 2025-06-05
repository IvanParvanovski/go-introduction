package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// func filterOldEmails(emails []email) {
// 	isOldChan := make(chan bool)

// 	for _, e := range emails {
// 		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.Local)) {
// 			isOldChan <- true
// 			continue
// 		}
// 		isOldChan <- false
// 	}

// 	isOld := <- isOldChan
// 	fmt.Println("email 1 is old:", isOld)
// 	isOld := <- isOldChan
// 	fmt.Println("email 2 is old:", isOld)
// 	isOld := <- isOldChan
// 	fmt.Println("email 3 is old:", isOld)
// }

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string)

	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println(email)
	}
}

func test(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	
	ch := addEmailsToQueue(emails)
	
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	
	fmt.Println("================")
}

func reader(id int, ch <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Hello ")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(4)

	go reader(1, ch, &wg)
	go reader(2, ch, &wg)
	go reader(3, ch, &wg)
	go reader(4, ch, &wg)

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()

}


