package main

import "fmt"

func main() {
	message, result := sumWithName("from main", 12039, 2, 3.5, 4.2, 5)

	fmt.Println(message, result)
}

// Return multiple values from the same function
func sumWithName(name string, numbers ...float32) (string, float32) {
	message := fmt.Sprintf("Sum of %s is: ", name)
	var result float32

	for _, number := range numbers {
		result += number
	}

	return message, result
}
