package packageone

import "fmt"

var PackageVar = "This is PackageVar in packageone"

func PrintMe(myVar, blockVar string) {
	fmt.Println(PackageVar)
	fmt.Println(myVar)
	fmt.Println(blockVar)
}
