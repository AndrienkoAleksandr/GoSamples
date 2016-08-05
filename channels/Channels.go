package main

import "fmt"

func main()  {
	message := make(chan string)
	fmt.Println(message)

	go func() {message <- "ping"} ()

	msg := <- message
	fmt.Println(msg)
}
