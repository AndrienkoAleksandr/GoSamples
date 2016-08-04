package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums)
	var sum int
	for _, elem := range nums {
		sum += elem
	}
	fmt.Println(" total :", sum)
	return sum
}

func main()  {
	sum(1, 2)
	sum(1, 2, 3)

	var res2 = []int{1, 2, 3, 4}
	sum(res2...)
}
