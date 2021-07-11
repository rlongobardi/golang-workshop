package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/* Consumer Producer problem, sending ten messages with a queue messages of N=5*/
func producer(from string, m chan string, size int) {
	i := 0
	for true {
		if len(m) < size {
			i++
			if checkDailyLimit(i, m) {
				break
			}
			message := "text number " + strconv.Itoa(i)
			fmt.Printf("Sending: %s from %s \n", message, from)
			m <- message
		} else {
			fmt.Println("Queue message full wait..")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func checkDailyLimit(i int, m chan string) bool {
	if i > 10 {
		fmt.Println("No sending anything else, sending signal!")
		m <- "done"
		return true
	}
	return false
}

func consumer(from string, m chan string) {
	for true {
		if len(m) >= 0 {
			receivedMessage := <-m
			fmt.Printf("Message received: %s consumed by %s \n", receivedMessage, from)
			if strings.EqualFold("done", receivedMessage) {
				fmt.Printf("Server has finished this communication...\n")
				break
			}
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
	go consumer("Consumer 1 ", message)

	go producer("Sender 1 ", message, sizeQueue)
	time.Sleep(1 * time.Second)

	fmt.Println("End producer & consumer")
}
