package main

import "fmt"

func main() {

	// Anonymous function (end with '()' and are automatically executed)
	func() {
		a := 1
		b := 2
		fmt.Println(a + b)
	}()

	// Assign an anonymous function to a variable
	myFunc := func() string {
		return "x"
	}

	// Print the memory reference (pointer) of the anonymous function
	fmt.Println(&myFunc)

	// Execute the anonymous function
	fmt.Println(myFunc())
}
