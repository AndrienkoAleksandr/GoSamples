package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res1 := <- c1:
		fmt.Println(res1)
	case <- time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

        select {
	case res1 := <- c2:
		fmt.Println(res1)
	case <- time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}
