package main

import (
	"fmt"
	"time"
)

func main() {
	i :=2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("tree")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	}

	currnetTime := time.Now()
	day := currnetTime.Weekday()
	fmt.Println("Today is ", day)
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}

	switch {
	case time.Hour < 12:
		fmt.Println("It's a before noon")
	default:
		fmt.Println("It's an after noon")
	}
}
