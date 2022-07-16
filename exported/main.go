package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{
		FirstName: "John",
		LastName:  "Doe",
		Salart:    1000,
		FullTime:  true,
	},
	{
		FirstName: "Jane",
		LastName:  "Doe",
		Salart:    2000,
		FullTime:  true,
	},
	{
		FirstName: "Joe",
		LastName:  "Doe",
		Salart:    3000,
		FullTime:  false,
	},
	{
		FirstName: "Jack",
		LastName:  "Doe",
		Salart:    4000,
		FullTime:  false,
	},
}

func main() {

	// Create a new Office struct
	myStaff := staff.Office{
		Employees: employees,
	}

	// myStaff.All()

	// fmt.Println(myStaff.All())
	staff.OverPaidLimit = 2000
	log.Println(myStaff.Overpaid())
	log.Println("Under paid staff:", myStaff.Underpaid())
	myStaff.NotVisible()

}

func myFunction() {
	log.Println("I am my function")
}
