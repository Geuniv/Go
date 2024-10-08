// 5-2
package main

import (
	"fmt"
)

// 자료형은 매개 변수 뒤에 붙여주고, 리턴값의 자료형도 맨 마지막에 지정함
func sum(a int, b int) int {
	return a + b
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
