package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready."

func main() {
	//seed the random generator

	rand.Seed(time.Now().UnixMicro())
	var firtsNumnber, secondNumber, subtraction, answer int

	generatorNumber(firtsNumnber, secondNumber, subtraction, answer)
}

func generatorNumber(firstNumber, secondNumber, subtraction, answer int) {
	firstNumber = rand.Intn(8) + 2
	secondNumber = rand.Intn(8) + 2
	subtraction = rand.Intn(8) + 2
	answer = firstNumber*secondNumber - subtraction
	reader := bufio.NewReader(os.Stdin)

	//dispaly a welcome/instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("-------------------------")
	fmt.Println("")

	fmt.Println("I'm thinking of a number between 1 and 10.Think and hit Enter to ready.")
	reader.ReadString('\n')
	//take them through the gmae
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divider the result by the number you  originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract the result from", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("The answer is", answer, ".")
}
