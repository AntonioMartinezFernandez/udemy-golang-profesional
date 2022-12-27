package main

import (
	"fmt"
)

func main() {

	fmt.Println("Conditionals -----------------------")

	var maxValue, inputValue int

	maxValue = 10

	fmt.Print("Enter value below 10: ")
	fmt.Scanln(&inputValue)

	if inputValue > maxValue {
		fmt.Println("Input too high")
	} else {
		fmt.Println("Correct input")
	}
}
