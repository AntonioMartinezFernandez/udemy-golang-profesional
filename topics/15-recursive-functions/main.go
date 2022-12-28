package main

import "fmt"

func main() {
	res := factorial(2)
	fmt.Println(res)
}

// Recursive functions are called by itself
func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}
