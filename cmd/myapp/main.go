package main

import (
	"fmt"
	"udemy-golang-profesional/pkg/challenges"
	"udemy-golang-profesional/pkg/variables"
)

func main() {
	fmt.Println("Hello", "World", "!")
	fmt.Println(variables.Integer())
	fmt.Println(variables.String())
	fmt.Printf("String: %s - Number: %d \n", "mystring", 5)
	text := fmt.Sprintf("text and variable %s", "string variable")
	fmt.Println(text)
	fmt.Println(variables.VarType(3))

	fmt.Println(challenges.Sum(1, 2))
	fmt.Println(challenges.Divide(10, 5))
	fmt.Println(challenges.Rest(8.5, 2))
	fmt.Println(challenges.Sellprice(10))
}
