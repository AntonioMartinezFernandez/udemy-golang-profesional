package main

import "fmt"

// Sum function
func sum(a, b int) int {
	return a + b
}

// Main function
func main() {

	// Variables
	var a, b int

	// Data input
	fmt.Print("First number: ")
	fmt.Scanln(&a)

	fmt.Print("Second number: ")
	fmt.Scanln(&b)

	// Calculate sum
	result := sum(a, b)

	// Data output
	fmt.Printf("%d + %d = %d", a, b, result)
}
