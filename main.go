package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)
// interface types

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

func main() {
	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Volvo",
		Model:         "XC()",
		Year:          2019,
	}

	fmt.Println("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
}
