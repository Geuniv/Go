// 4-7
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5, 10)

	fmt.Println(len(a))
	fmt.Println(cap(a))
}
