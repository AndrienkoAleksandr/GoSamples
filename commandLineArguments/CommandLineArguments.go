package main

import (
	"os"
	"fmt"
)

func main()  {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
