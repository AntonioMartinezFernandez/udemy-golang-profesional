package main

import (
	"fmt"
	"math"
)

// Divide function
func divide(a, b float64) (float64, float64) {
	result := a / b
	rest := math.Mod(a, b)
	return result, rest
}

// Main function
func main() {

	// Variables
	var a, b float64

	// Data input
	fmt.Print("First number: ")
	fmt.Scanln(&a)

	fmt.Print("Second number: ")
	fmt.Scanln(&b)

	// Calculate division
	result, rest := divide(a, b)

	// Data output
	fmt.Printf("%f / %f = %f \n", a, b, result)
	fmt.Printf("Rest: %f", rest)
}
