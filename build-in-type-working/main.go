package main

import "fmt"

// aggregate types (pointer, struct)

func main() {

	x := 10

	myFirstPointer := &x

	fmt.Println("x is:", x)
	fmt.Println("myFirstPointer is:", myFirstPointer)

	changeValueOfPointer(&x)

	fmt.Println("x is after func called:", x)

}

func changeValueOfPointer(myPointer *int) {
	*myPointer = 15
}
