package main

import (
	"fmt"
)

func main() {
	names := [...]string{"Stromae", "Miles", "Kurt"}

	fmt.Println("For (normal) -----------------------")

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	fmt.Println("For each -----------------------")
	for index, element := range names {
		fmt.Println(index, element)
	}

}
