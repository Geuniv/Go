// 4-11
package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
