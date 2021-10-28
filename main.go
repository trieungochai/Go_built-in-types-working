package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	z := sumMany(1, 2, 3, 4, -5, 21, -100)
	fmt.Println(z)
}

func sumMany(nums ...int) int {
	total := 0
	for _, x := range nums {
		total = total + x
	}

	return total
}
