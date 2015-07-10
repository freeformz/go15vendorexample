package c

import "a"
import "fmt"

var C = 3

func init() {
	fmt.Println("Init package c")
	fmt.Println("c: C ==", C)
	fmt.Println("c: a.A ==", a.A)
	fmt.Println("c: a.Inc() ==", a.Inc())
}
