package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("1st element in array list is", myStrings[0])
	fmt.Println("2st element in array list is", myStrings[1])
	fmt.Println("3rd element in array list is", myStrings[2])
}
