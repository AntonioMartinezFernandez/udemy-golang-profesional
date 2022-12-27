package main

import (
	"fmt"
)

// Divide function
func taxes(a float32) float32 {
	tax := a * 0.21
	result := a + tax
	return result
}

// Main function
func main() {

	// Variables
	var a float32

	// Data input
	fmt.Print("Price: ")
	fmt.Scanln(&a)

	// Calculate division
	result := taxes(a)

	// Data output
	fmt.Printf("Price with taxes: %f \n", result)
}
