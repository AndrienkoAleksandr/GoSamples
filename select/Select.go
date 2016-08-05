package main

import (
	"time"
	"fmt"
)

func main()  {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	sortedMessages := make([]string, 2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
			sortedMessages[0] = msg1
		case msg2 := <- c2:
			sortedMessages[1] = msg2
		}
	}
	fmt.Println("Received:", sortedMessages)
}
