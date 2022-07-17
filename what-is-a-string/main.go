package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println()
	name := "Hello World"

	fmt.Println(name)

	fmt.Println("Bytes")

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRunes\tString")

	for i, c := range name {
		fmt.Println(i, "\t", c, "\t", string(c))
	}

	fmt.Println()

	h := "Hello"
	w := "World"

	myString := fmt.Sprintf("%s %s", h, w)
	fmt.Println("myString:", myString)
	fmt.Println()

	//using string builder

	sb := new(strings.Builder)
	sb.WriteString(h)
	sb.WriteString(" ")
	sb.WriteString(w)
	fmt.Println("sb:", sb.String())

	fmt.Println()

	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	fmt.Println("Getting a substring")
	fmt.Println(name[2:5])

}
