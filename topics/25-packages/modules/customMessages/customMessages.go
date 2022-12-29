package customMessages

import "fmt"

func PrintHello() {
	fmt.Println("Hello from messages package")
}

// If function name starts with uppercase, the function is public
func PrintPrivateMessage() {
	fmt.Println(privateMessage())
}

// If function name starts with lowercase, the function is private
func privateMessage() string {
	return "This is a private message"
}
