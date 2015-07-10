package b

import "a"
import "fmt"

var B = 2

func init() {
	fmt.Println("Init package b")
	fmt.Println("b: a.A ==", a.A)
	fmt.Println("b: a.Inc() ==", a.Inc())
}
