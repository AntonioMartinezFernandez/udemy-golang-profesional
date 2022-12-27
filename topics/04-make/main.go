package main

import (
	"fmt"
)

func main() {

	fmt.Println("Make -----------------------")

	numbers := make([]int, 2, 2)

	numbers[0] = 12
	numbers[1] = 25

	fmt.Println(numbers)
	fmt.Printf("Len: %v, Cap: %v, MemRef: %p  \n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 82, 45)

	fmt.Println(numbers)
	fmt.Printf("Len: %v, Cap: %v, MemRef: %p  \n", len(numbers), cap(numbers), numbers)
}
