package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Coffee"
	coffees[2] = "Tea"
	coffees[3] = "Coffee with milk"
	coffees[4] = "Tea with milk"
	coffees[5] = "Coffee with sugar"

	fmt.Println("Menu")
	fmt.Println("----------------")
	fmt.Println("1. Capuchino")
	fmt.Println("2. Latte")
	fmt.Println("3.Americano")
	fmt.Println("4. Mocha")
	fmt.Println("5. Machiato")
	fmt.Println("6. Expresso")
	fmt.Println("7. Quit")
	fmt.Println("Q- Quit the program")

	fmt.Println("Press any key on the keyboard. Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}
		i, _ := strconv.Atoi(string(char))

		// _, ok := coffees[i]
		// if ok {

		// 	t := fmt.Sprintf("You choose %s", coffees[i])
		// 	fmt.Println(t)
		// }

		if _, ok := coffees[i]; ok {
			t := fmt.Sprintf("You choose %s", coffees[i])
			fmt.Println(t)
		}

		if key == keyboard.KeyEsc {
			fmt.Printf("Program exiting...")
			break
		}

	}
}
