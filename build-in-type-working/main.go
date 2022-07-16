package main

import (
	"fmt"
	"sort"
)

// aggregate types (pointer, slices)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "bird")
	animals = append(animals, "fish")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(x, i)
	}

	fmt.Println(animals[1:3])

	fmt.Println(animals[:3])

	fmt.Println(animals[3:])

	fmt.Println("The slice is:", len(animals), "long")

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)

}
func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] //1 2 3 4 5
	a[len(a)-1] = ""   //1 2 3 4 5
	a = a[:len(a)-1]   //1 2 3 4
	return a
}
