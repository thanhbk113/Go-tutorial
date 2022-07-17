package main

import (
	"fmt"
	"strings"
)

func main() {

	myString := "This is a clear example of a why we search in one case only"

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		println("Found")
	} else {
		println("Not Found")
	}

	//other case function
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))

}
