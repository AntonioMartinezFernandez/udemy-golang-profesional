package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...]string{"Stromae", "Miles", "Kurt"}

	fmt.Println("ToLower & ToUpper -----------------------")
	for _, element := range names {
		fmt.Println(strings.ToLower(element))
		fmt.Println(strings.ToUpper(element))
	}

	fmt.Println("Replace -----------------------")
	repString := "Blue light blue"

	fmt.Println(strings.Replace(repString, " ", "_", 1))
	fmt.Println(strings.ReplaceAll(repString, " ", "_"))

	fmt.Println("Split & Join -----------------------")
	myWord := "abcba"
	reverseWord := reverseString(myWord)
	fmt.Println(myWord, "is palyndrome? ", strings.ToUpper(myWord) == reverseWord)

}

func reverseString(myString string) string {
	// to upper case
	myString = strings.ToUpper(myString)

	// Split method
	strArray := strings.Split(myString, "")

	// Reverse string array
	var revStrArray []string
	for i := len(strArray); i > 0; i-- {
		revStrArray = append(revStrArray, strArray[i-1])
	}

	// Join method
	return strings.Join(revStrArray, "")
}
