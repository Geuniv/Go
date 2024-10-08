// 5-9
package main

import (
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

func main() {
	var hello func(a int, b int) int = sum
	world := sum

	fmt.Println(hello(1, 2))
	fmt.Println(world(1, 2))
}
