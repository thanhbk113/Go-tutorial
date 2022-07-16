package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// aggregate types (pointer, slices,maps,function,go routine)

var keyPressChanel chan rune

func main() {

	keyPressChanel = make(chan rune)
	go listenForKeyPress() // start listening for key presses in a go routine
	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer keyboard.Close()

	for {
		char, _, _ := keyboard.GetKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChanel <- char
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChanel
		fmt.Println("You pressed:", string(key))
	}
}

// func doSomeThing(s string) {
// 	until := 0
// 	for {
// 		fmt.Println("s is ", s)
// 		until++
// 		if until >= 5 {
// 			break
// 		}
// 	}
// }
