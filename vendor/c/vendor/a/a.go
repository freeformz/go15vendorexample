package a

import "fmt"

var A = 100

func Inc() int {
	A++
	return A
}
func init() {
	fmt.Println("Init c/vendor/a")
	fmt.Println("c/vendor/a: A ==", A)
}
