package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Calculator -----------------------")

	var operation string

	fmt.Println("Operation (a+b+c): ")
	fmt.Scanln(&operation)

	fmt.Println("Result: ", sum(operation))
}

func sum(operation string) int {
	values := strings.Split(operation, "+")
	result := 0

	for i := len(values); i > 0; i-- {
		val, err := strconv.Atoi(values[i-1])

		// Error handler
		if err != nil {
			fmt.Println("Error", err)
			return 0
		}

		result += val
	}

	return result
}
