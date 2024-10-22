// 6-1
package main

import (
	"fmt"
)

func main() {
	var numPtr *int = new(int)

	*numPtr = 1

	fmt.Println(*numPtr)
}
