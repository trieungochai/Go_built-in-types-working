package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)
// interface types

func main() {
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "frog")
	animals = append(animals, "fish")
	animals = append(animals, "pig")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}

	fmt.Println("Element 0 is", animals[0])
	fmt.Println("First two element are", animals[0:2])
	fmt.Println("The slice is", len(animals), "elements long")
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)
}
