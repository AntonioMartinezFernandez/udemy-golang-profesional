package main

import (
	"fmt"
)

func main() {

	fmt.Println("Relational Operators -----------------------")
	var a, b int
	a = 1
	b = 2

	fmt.Printf("a: %d, b: %d, a == b? %t \n", a, b, a == b)
	fmt.Printf("a: %d, b: %d, a != b? %t \n", a, b, a != b)
	fmt.Printf("a: %d, b: %d, a < b? %t \n", a, b, a < b)
	fmt.Printf("a: %d, b: %d, a > b? %t \n", a, b, a > b)
	fmt.Printf("a: %d, b: %d, a <= b? %t \n", a, b, a <= b)
	fmt.Printf("a: %d, b: %d, a >= b? %t \n", a, b, a >= b)

	fmt.Println("Logic Operators -----------------------")
	var c, d bool
	c = true
	d = false

	fmt.Printf("c: %t, !c? %t \n", c, !c)
	fmt.Printf("c: %t, d: %t, c && d? %t \n", c, d, c && d)
	fmt.Printf("c: %t, d: %t, c || d? %t \n", c, d, c || d)
}
