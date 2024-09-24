// 4-24
package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"Hello": 10, "world": 20}

	delete(a, "world")

	fmt.Println(a)
}
