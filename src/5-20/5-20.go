// 5-20
package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2}

	for i := 0; i < 3; i++ { // 배열을 벗어난 접근을 함
		fmt.Println(a[i])
	}
}
