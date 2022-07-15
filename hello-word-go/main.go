package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		userInput, _ := reader.ReadString('\n')

		if userInput == "quit\r\n" || userInput == "quit\n" {
			break
		}

		fmt.Println(doctor.Response(userInput))

	}

}
