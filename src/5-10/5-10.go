// 5-10
package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := []func(int, int) int{sum, diff}
	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](1, 2))
}
