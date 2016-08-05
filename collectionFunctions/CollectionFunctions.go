package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, elem := range vs {
		if elem == t {
			return i
		}
	}
	return -1;
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string)bool) bool  {
	for _, elem := range vs {
		if f(elem) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, elem := range vs {
		if !f(elem) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(s string) bool) []string {
	var results = make([]string, 0)
	for _, elem := range vs {
		if f(elem) {
			results = append(results, elem)
		}
	}
	return results;
}

func Map(vs []string, f func(s string) string) []string  {
	var res = make([]string, len(vs))
	for i, elem := range vs {
		res[i] = f(elem)
	}
	return res
}

func main()  {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v , "p")
	}))

	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "p")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
