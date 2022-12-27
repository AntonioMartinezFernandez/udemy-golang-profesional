package main

import (
	"fmt"
)

func main() {

	fmt.Println("Arrays -----------------------")

	var myArray [5]int
	myArray[0] = 1
	myArray[1] = 2
	fmt.Println(myArray[0])
	fmt.Println(myArray[2])
	fmt.Println(myArray)

	myStringsArray := [...]string{"myFirstString", "mySecondString"}
	myStringsArray[1] = "myModifiedSecondString"
	fmt.Println(myStringsArray, len(myStringsArray))

	var otherStringArray [5]string
	otherStringArray = [5]string{
		0: "hello",
		1: "bye",
	}
	otherStringArray[2] = "another"
	fmt.Println(otherStringArray, len(otherStringArray))

	subOtherArray := otherStringArray[0:2]
	fmt.Println(subOtherArray, len(subOtherArray))

	fmt.Println("Slices ------------------------")

	numbersSlice := []int{1, 2, 3}
	fmt.Println(numbersSlice, len(numbersSlice))
	numbersSlice = append(numbersSlice, 4)
	numbersSlice = append(numbersSlice, 5)
	fmt.Println(numbersSlice, len(numbersSlice))

	subNumbersSlice := numbersSlice[0:2]         // Extract an array from numbersSlice, 2 elements from position 0
	subNumbersSlice[1] = subNumbersSlice[1] * 50 // !Changes subNumbersSlice AND numbersSlice

	fmt.Println(subNumbersSlice)
	fmt.Println(numbersSlice)

	fmt.Println("Slices 2 ------------------------")

	months := []string{"January", "February"}
	fmt.Printf("Len: %v, Cap: %v, MemRef: %p \n", len(months), cap(months), months)

	months = append(months, "March", "April")
	fmt.Printf("Len: %v, Cap: %v, MemRef: %p \n", len(months), cap(months), months)
}
