package main

import "fmt"

func main() {
	var str string = "My string"

	fmt.Printf("main memref: %p \n", &str)

	fmt.Println("string from main: ", str)

	modifyValue(str)

	fmt.Println("string after value-modifier function: ", str)

	modifyMemref(&str)

	fmt.Println("string after memref-modifier function: ", str)
}

func modifyValue(str string) {
	fmt.Printf("value-modifier memref: %p \n", &str)
	str = "Modified from value-modifier"
}

func modifyMemref(str *string) {
	fmt.Printf("memref-modifier memref: %p \n", str)
	*str = "Modified from memref-modifier"
}
