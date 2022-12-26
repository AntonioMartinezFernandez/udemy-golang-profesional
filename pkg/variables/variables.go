package variables

import "fmt"

func Integer() int {
	integer := 2
	return integer
}

func String() string {
	string := "string"
	return string
}

func VarType(Var int) string {
	return fmt.Sprintf("%T", Var)
}
