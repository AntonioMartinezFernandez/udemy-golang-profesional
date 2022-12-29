package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Imported package -----------------------")

	fmt.Println(Sum(1, 2))
	fmt.Println(Divide(10, 5))
	fmt.Println(Rest(8.5, 2))
	fmt.Println(Sellprice(10))
}

func Sum(a, b int) int {
	return a + b
}

func Divide(a, b float32) float32 {
	return a / b
}

func Rest(a, b float64) float64 {
	return math.Mod(a, b)
}

func Sellprice(a float32) float32 {
	igv := a * 0.18
	return a + igv
}
