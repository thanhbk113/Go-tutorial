package main

import (
	"fmt"
	"strings"
)

func main() {
	// courses := []string{"Learn Go to Beginner", "Learn Go to Intermediate", "Learn Go to Advanced"}

	// for _, course := range courses {

	// 	if strings.Contains(course, "Go") {
	// 		println("Go is found in", course, "and index is", strings.Index(course, "Go"))
	// 	}
	// }

	newString := "Go is a great programming language.Go for it!"

	// println(strings.HasPrefix(newString, "Go"))
	// println(strings.HasSuffix(newString, "Python"))
	// println(strings.HasSuffix(newString, "!"))
	// println(strings.Count(newString, "Go"))
	// println(strings.Count(newString, "Python"))
	// println(strings.Index(newString, "Python"))

	// newString = strings.Replace(newString, "Go", "Golang", -1)
	newString = strings.ReplaceAll(newString, "Go", "Golang")

	println(newString)

	if "A" > "B" {
		println("A is greater than B")
	} else {
		println("B is greater than A")
	}

	badEmail := " me@here.com "
	badEmail = strings.TrimSpace(badEmail)

	fmt.Printf("=%s=", badEmail)
}
