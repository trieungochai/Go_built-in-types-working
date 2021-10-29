package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	go doSomething("Hello World!!!")

	for {
	}
}

func doSomething(s string) {
	for {
		fmt.Println("s is", s)
	}
}
