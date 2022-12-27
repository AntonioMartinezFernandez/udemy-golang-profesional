package main

import (
	"fmt"
)

func main() {

	fmt.Println("Conditionals 1 -----------------------")

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

	fmt.Println("Conditionals 2 -----------------------")

	// Variables declared in the 'if' scope
	if name, age := "Stromae", 38; age < 40 {
		fmt.Printf("Hi %s! You're young! \n", name)
	}

	groups := make(map[string]int)

	groups["Blur"] = 1991

	if year, exists := groups["Blur"]; exists {
		fmt.Println("First album: ", year)
	} else {
		fmt.Println("Group doesn't exists")
	}

	if year, exists := groups["U2"]; exists {
		fmt.Println("First album: ", year)
	} else {
		fmt.Println("Group doesn't exists")
	}

}
