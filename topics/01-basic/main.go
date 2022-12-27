package main

import (
	"fmt"
	"udemy-golang-profesional/pkg/variables"
)

func main() {

	fmt.Println("Basic -----------------------")

	fmt.Println("Hello", "World", "!")
	fmt.Println(variables.Integer())
	fmt.Println(variables.String())
	fmt.Printf("String: %s - Number: %d \n", "mystring", 5)
	text := fmt.Sprintf("text and variable %s", "string variable")
	fmt.Println(text)
	fmt.Println(variables.VarType(3))
}
