package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)
// interface types

type Animals struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animals) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animals) HowmanyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	var dog Animals
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowmanyLegs()

	cat := Animals{
		Name:         "cat",
		Sound:        "moew",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowmanyLegs()
}
