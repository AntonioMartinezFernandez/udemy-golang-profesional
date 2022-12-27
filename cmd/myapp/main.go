// fmt verbs -> https://pkg.go.dev/fmt

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Conditionals -----------------------")

	var minValue, maxValue, inputValue float64

	minValue = 0
	maxValue = 10

	fmt.Print("Enter value between 0 and 10: ")
	fmt.Scanln(&inputValue)

	if inputValue > maxValue || inputValue < minValue {
		fmt.Println("Invalid input")
	} else {
		fmt.Println("Correct input")
	}
}
