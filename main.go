package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	z := addTwoNumber(1, 2)
	fmt.Print(z)
}

func addTwoNumber(x, y int) int {
	return x + y
}
