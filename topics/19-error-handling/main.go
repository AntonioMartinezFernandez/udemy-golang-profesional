package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(2, 0)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}

	result, err = divide(10, 2)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("Division by 0")
	}

	result := a / b

	return result, nil
}
