// 5-6
package main

import (
	"fmt"
)

// 가변 인자 사용
// 함수의 매개변수 개수가 정해져 있지 않고 유동적으로 변하는 형태를 가변인자라고함
func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}
	return total
}

func main() {
	r := sum(1, 2, 3, 4, 5)

	fmt.Println(r)
}
