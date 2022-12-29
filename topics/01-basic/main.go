package main

import (
	"fmt"
)

func main() {

	fmt.Println("Basic -----------------------")

	fmt.Println("Hello", "World", "!")
	fmt.Println(Integer())
	fmt.Println(String())
	fmt.Printf("String: %s - Number: %d \n", "mystring", 5)
	text := fmt.Sprintf("text and variable %s", "string variable")
	fmt.Println(text)
	fmt.Println(VarType(3))
}

func Integer() int {
	integer := 2
	return integer
}

func String() string {
	string := "string"
	return string
}

func VarType(Var int) string {
	return fmt.Sprintf("%T", Var)
}
