package main

import "myapp/packageone"

var myVar = "This is a global variable"

func main() {

	var blockVar = "This is a block variable"

	packageone.PrintMe(myVar, blockVar)

	//in the main function,print out the value of myVar
	//blockvar, and Package on one line using the PrintMe
	//function in packageone

}
