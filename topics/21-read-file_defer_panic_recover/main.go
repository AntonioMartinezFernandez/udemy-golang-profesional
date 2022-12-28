package main

import (
	"fmt"
	"os"
)

// FIRST APROXIMATION -BASIC STYLE-

// func main() {
// 	file, err := os.Open("./topics/21-read-file_defer_panic_recover/file.txt")

// 	if err != nil {
// 		fmt.Println("Error opening file.txt", err)
// 	} else {
// 		content := make([]byte, 256)
// 		long, _ := file.Read(content)
// 		fmt.Println(content)
// 		fileContent := string(content[0:long])
// 		fmt.Println(fileContent)
// 	}

// 	file.Close()
// }

// SECOND APROXIMATION -GO STYLE AND DEFER-

// func main() {
// 	if file, err := os.Open("./topics/21-read-file_defer_panic_recover/file.txt"); err != nil {
// 		fmt.Println("Error opening file.txt", err)
// 	} else {
// 		// defer execute the function (anonymous in this case) at the end of the scope
// 		defer func() {
// 			fmt.Println("Closing file...")
// 			file.Close()
// 		}()

// 		content := make([]byte, 256)
// 		long, _ := file.Read(content)
// 		fmt.Println(content)
// 		fileContent := string(content[0:long])
// 		fmt.Println(fileContent)
// 	}
// }

// THIRD APROXIMATION -GO STYLE, DEFER AND PANIC-

// func main() {
// 	if file, err := os.Open("./topics/21-read-file_defer_panic_recover/fil.txt"); err != nil {
// 		// panic -> stop the execution in a hard way
// 		panic("Error opening file.txt")
// 	} else {
// 		// defer execute the function (anonymous in this case) at the end of the scope
// 		defer func() {
// 			fmt.Println("Closing file...")
// 			file.Close()
// 		}()

// 		content := make([]byte, 256)
// 		long, _ := file.Read(content)
// 		fmt.Println(content)
// 		fileContent := string(content[0:long])
// 		fmt.Println(fileContent)
// 	}
// }

// FOURTH APROXIMATION -GO STYLE, DEFER, PANIC AND RECOVER-

func main() {
	defer func() {
		// recover -> control exceptions watching the error variable
		if err := recover(); err != nil {
			fmt.Println("Incorrect execution end -message from recover-")
		}
	}()

	if file, err := os.Open("./topics/21-read-file_defer_panic_recover/file.txt"); err != nil {
		// panic -> stop the execution in a hard way
		panic("Error opening file.txt")
	} else {
		// defer execute the function (anonymous in this case) at the end of the scope
		defer func() {
			fmt.Println("Closing file...")
			file.Close()
		}()

		content := make([]byte, 256)
		long, _ := file.Read(content)
		fmt.Println(content)
		fileContent := string(content[0:long])
		fmt.Println(fileContent)
	}
}
