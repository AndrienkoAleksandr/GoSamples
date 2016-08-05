package main

import (
	"time"
	"fmt"
)

func main()  {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++  {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for request := range requests {
		<- limiter
		fmt.Println("request:", request, time.Now())
	}



	burstyLimiters := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiters <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiters <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++  {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range  burstyRequests {
		<- burstyLimiters
		fmt.Println("request:", req, time.Now())
	}
}
