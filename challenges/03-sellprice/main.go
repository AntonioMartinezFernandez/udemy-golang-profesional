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
	totalPrice := taxes(a)

	// Data output
	// fmt verbs -> https://pkg.go.dev/fmt
	result := fmt.Sprintf("%.2f", totalPrice)
	fmt.Printf("Price with taxes: %s \n", result)
}
