package main

import "fmt"

//expression

func main() {

	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old.Right handed : %t\n", name, age, rightHanded)

	fmt.Println()

	ageInTenYears := age + 10 // expression

	fmt.Printf("%s will be %d years old in 10 years.\n", name, ageInTenYears)

	isATeenager := ageInTenYears >= 13 // expression

	fmt.Printf("%s is a teenager : %t\n", name, isATeenager)
}
