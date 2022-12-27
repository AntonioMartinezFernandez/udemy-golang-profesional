package main

import (
	"fmt"
)

func main() {

	fmt.Println("Conditionals Switch -----------------------")

	var letter string

	fmt.Print("Write a letter: ")
	fmt.Scanln(&letter)

	switch {
	case letter == "a" || letter == "A":
		fmt.Println("is vowel (a)")
		break

	case letter == "e" || letter == "E":
		fmt.Println("is vowel (e)")
		break

	case letter == "i" || letter == "I":
		fmt.Println("is vowel (i)")
		break

	case letter == "o" || letter == "O":
		fmt.Println("is vowel (o)")
		break

	case letter == "u" || letter == "U":
		fmt.Println("is vowel (u)")
		break

	default:
		fmt.Println("is consonant")
	}

	switch letter {
	case "a", "A":
		fmt.Println("is vowel (a)")
		break

	case "e", "E":
		fmt.Println("is vowel (e)")
		break

	case "i", "I":
		fmt.Println("is vowel (i)")
		break

	case "o", "O":
		fmt.Println("is vowel (o)")
		break

	case "u", "U":
		fmt.Println("is vowel (u)")
		break

	default:
		fmt.Println("is consonant")
	}
}
