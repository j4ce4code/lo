package lo

import "fmt"

func PrintVersion() {
	fmt.Println("0.1.4")
}

func Ternary(condition bool, this any, that any) any {
	if condition {
		return this
	}
	return that
}
