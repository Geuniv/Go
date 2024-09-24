// 4-19
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a[:2]
	b[0] = 9

	fmt.Println(a)
	fmt.Println(b)
}
