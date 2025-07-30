package main

import (
	"bufio"
	"ethos/fmt"
	"os"
)

func main() {
	fmt.Print("Enter some text: ")
	reader := bufio.NewReader(os.Stdin)

	// Read until a newline character
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered:", input)
}