// 5-5
package main

import (
	"fmt"
)

// 값을 여러 개 리턴할 때도 리턴값에 이름을 정할 수 있음
func SumAndDiff(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)
}
