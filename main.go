package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat32 float32
var myFloat64 float64

// aggregate types (array, struct)
// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	myInt = 10
	myUint = 20
	myFloat32 = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat32, myFloat64)

	myString := "trieungochai"
	log.Println(myString)
	myString = "cai"

	var myBool = true
	myBool = false
	log.Println(myBool)
}
