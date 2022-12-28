package main

import "fmt"

// Defined variables and constants out of the functions scope are globals
var message string = "global variable"

const constant string = "you can't modify constants"

func main() {
	// defer -> move the execution to the end of the scope
	defer modifyAndPrint()

	fmt.Println(message)
	fmt.Println(constant)

	printer()
}

func printer() {
	fmt.Println(message)
	fmt.Println(constant)
}

func modifyAndPrint() {
	message = "modified global variable"
	fmt.Println(message)
}
