package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	//execute a loop while i > 100

	i := 1000
	for i > 100 {
		//get a random number between 1 and 1001
		i = rand.Intn(1001) + 1
		fmt.Println(i)

		if i > 100 {
			fmt.Println("i is ", i)
		}
	}

	fmt.Println("Got", i, "and broke out of the loop")
}
