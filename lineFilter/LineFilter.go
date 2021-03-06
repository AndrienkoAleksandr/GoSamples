package main

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := scanner.Text()
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
