// 4-2
package main

import (
	"fmt"
)

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
