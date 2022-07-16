package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {

	jack := Employee{
		Name:     "Jack Smith",
		Age:      2,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is 30 or older")
		} else {
			fmt.Println(x.Name, "is under 30")
		}

		if x.Age > 30 && x.Salary > 5000 {
			fmt.Println(x.Name, " make more than 50000 and is over 30")
		} else {
			fmt.Println(x.Name, " make less than 50000 and is under 30")

		}
	}
}
