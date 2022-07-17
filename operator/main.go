package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 12
	b := 0

	c, err := divideTwoNumbers(a, b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("c:", c)
	}

}

func divideTwoNumbers(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
