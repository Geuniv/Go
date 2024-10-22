// 6-3
package main

import (
	"fmt"
)

func hello(n int) {
	n = 2
}

func main() {
	var n int = 1

	hello(n)

	fmt.Println(n)
}
