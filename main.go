package main

import (
	"a"
	"b"
	"c"
	"fmt"
)

func main() {
	fmt.Println("main.main()")
	fmt.Println("main: a.A ==", a.A)
	fmt.Println("main: a.Inc() ==", a.Inc())
	fmt.Println("main: b.B ==", b.B)
	fmt.Println("main: c.C ==", c.C)
}
