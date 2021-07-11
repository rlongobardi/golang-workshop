package main

import (
	"fmt"
	"strconv"
	"time"
)

func producer(m chan string, size int) {
	for i := 0; i < 10; i++ {
		if len(m) < size {
			message := "text number " + strconv.Itoa(i)
			fmt.Printf("Sending: %s\n", message)
			m <- message
		} else {
			fmt.Println("Queue message full wait..")
			time.Sleep(100 * time.Millisecond)
		}

	}
}

func consumer(from string, m chan string) {
	for i := 0; i < 10; i++ {
		if len(m) > 0 {
			fmt.Printf("Message received: %s consumed by %s \n", <-m, from)
		} else {
			fmt.Println("Queue messages is empty")
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func main() {
	sizeQueue := 5
	message := make(chan string, sizeQueue)
	fmt.Println("Start producer & consumer")
	go producer(message, sizeQueue)
	go consumer("Consumer XXX-111", message)
	go consumer("Consumer XXX-222", message)
	fmt.Println("End producer & consumer")
	time.Sleep(time.Second)
	<-message // deadlock

}
