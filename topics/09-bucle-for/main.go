package main

import (
	"fmt"
)

func main() {

	fmt.Println("For (infinite) -----------------------")
	counter := 0
	for {

		counter++

		if counter <= 5 {
			fmt.Println("continue infinite bucle", counter)
			continue
		}

		if counter >= 10 {
			fmt.Println("break infinite bucle", counter)
			break
		}

		fmt.Println("Infinite hello")
	}

	fmt.Println("For (while type) -----------------------")
	number := 10000
	count := 0

	fmt.Println("Number", number)

	for number > 0 {
		number /= 10
		count++
	}

	fmt.Println("Digits: ", count)

	fmt.Println("For (normal) -----------------------")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
