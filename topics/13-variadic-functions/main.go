package main

import "fmt"

func main() {
	fmt.Print(sum(12039, 2, 3.5, 4.2, 5))
}

// Variadic functions receive an indetermined quantity of parameters as array
func sum(numbers ...float32) float32 {
	fmt.Println(numbers)

	var result float32

	for _, number := range numbers {
		result += number
	}

	return result
}
