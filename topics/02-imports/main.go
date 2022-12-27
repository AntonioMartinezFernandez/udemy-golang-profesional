package main

import (
	"fmt"
	"udemy-golang-profesional/pkg/challenges"
)

func main() {

	fmt.Println("Imported package -----------------------")

	fmt.Println(challenges.Sum(1, 2))
	fmt.Println(challenges.Divide(10, 5))
	fmt.Println(challenges.Rest(8.5, 2))
	fmt.Println(challenges.Sellprice(10))
}
