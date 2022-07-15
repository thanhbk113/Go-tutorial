package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("What is your name? ")
	user.Age = readInt("How old are you? ")
	user.FavoriteNumber = readFloat("What is your favorite number? ")
	user.OwnsADog = readBoolean("Do you own a dog? ")
	prompt()

	fmt.Printf("Your name is %s, and you are %d years old.Your favorite number is %.4f.Owns a dog: %t", user.UserName, user.Age, user.FavoriteNumber, user.OwnsADog)

}

func prompt() {
	fmt.Print("--> ")
}

func readString(s string) string {

	for {
		fmt.Print(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "/r/n", "", -1)
		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a valid string")
		}

	}
}

func readInt(s string) int {
	for {

		fmt.Print(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {

		fmt.Print(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBoolean(s string) bool {
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Error opening keyboard")
	}

	defer func() {
		keyboard.Close()
	}()

	for {

		fmt.Println(s)

		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if char == 'n' || char == 'N' {
			return false
		} else if char == 'y' || char == 'Y' {
			return true
		}

	}
}
