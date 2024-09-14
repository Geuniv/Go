// 4-6
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a // 배열을 대입하면 배열 전체가 복사됨

	fmt.Println(a)
	fmt.Println(b)
}
