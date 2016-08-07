package main

import (
	"os"
	"fmt"
	"strings"
)

func main()  {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	for i, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(i, " ", pair[0])
	}
}
