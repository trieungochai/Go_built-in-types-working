package main

import (
	"fmt"
)

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	x := 10
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)
	fmt.Println("myFirstPointer is", myFirstPointer)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
