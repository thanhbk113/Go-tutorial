package main

import "fmt"

// aggregate types (pointer, slices,maps,function)

type Animal struct {
	name         string
	sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.name, a.sound)
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs\n", a.name, a.NumberOfLegs)
}

func main() {

	moo := Animal{"cow", "moo", 4}
	moo.Says()
	moo.HowManyLegs()
	bob := Animal{"bird", "tweet", 2}
	bob.Says()
	bob.HowManyLegs()

}

// func addTwoNumber(a, b int) int {
// 	return a + b
// }

// func sumMany(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }
