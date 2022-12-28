package main

import (
	"fmt"
	"strings"
)

func main() {
	closure := repeat(5)

	resultOne := closure("repeat this.")
	fmt.Println(resultOne)

	resultTwo := closure("another string.")
	fmt.Println(resultTwo)
}

// Closure: function return another function
func repeat(times int) func(str string) string {
	return func(str string) string {
		return strings.Repeat(str, times)
	}
}
