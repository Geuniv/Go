// 5-13
package main

import (
	"fmt"
)

func main() { // 함수 안에서
	sum := func(a, b int) int { // 익명 함수를 선언 및 정의
		return a + b
	}

	r := sum(1, 2)
	fmt.Println(r)
}
