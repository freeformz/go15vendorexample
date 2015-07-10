package a

import "fmt"

var A = 1

func Inc() int {
	A++
	return A
}
func init() {
	fmt.Println("Init package a")
	fmt.Println("a: A ==", A)
}
