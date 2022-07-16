package main

import (
	"bufio"
	"fmt"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something: ")

	for i := 0; i < 5; i++ {
		fmt.Println("--> ")

		text, _ := reader.ReadString('\n')

		ch <- text

		time.Sleep(time.Second)
	}

}
