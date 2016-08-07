package main

import (
	"os"
	"fmt"
)

func main()  {
	defer fmt.Println("!")

	os.Exit(3)
}
