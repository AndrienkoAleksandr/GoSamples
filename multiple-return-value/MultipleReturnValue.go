package main

import "fmt"

func vals() (int, int, int) {
	return 3, 7, 1
}

func main()  {
	a, b, c := vals()
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)

	fmt.Println()

	_, d, f := vals()
	fmt.Print(d)
	fmt.Print(f)
}
