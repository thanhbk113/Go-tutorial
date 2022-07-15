package main

import "fmt"

// aggregate types (array, struct)

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
		Make:          "Ford",
		Model:         "Fusion",
		Year:          2015,
	}

	// myCar.NumberOfTires = 4
	// myCar.Luxury = false
	// myCar.BucketSeats = true
	// myCar.Make = "Ford"
	// myCar.Model = "Fusion"
	// myCar.Year = 2018

	fmt.Printf("My car is a %s %s %d\n", myCar.Make, myCar.Model, myCar.Year)

}
