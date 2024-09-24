// 4-12
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	var b []int

	b = a
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
