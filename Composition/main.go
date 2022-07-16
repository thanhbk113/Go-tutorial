package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	println("Number of wheels:", v.NumberOfWheels)
	println("Number of passengers:", v.NumberOfPassengers)
}

func (c Car) show() {
	println("Make:", c.Make)
	println("Model:", c.Model)
	println("Year:", c.Year)
	println("Is Electric:", c.IsElectric)
	println("Is Hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {

	suv := Vehicle{4, 5}
	suv.showDetails()

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2020,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	volvoXC90.show()

	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2020,
		IsElectric: true,
		IsHybrid:   true,
		Vehicle:    suv,
	}

	teslaModelX.show()

}
