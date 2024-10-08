// 5-4
package main

import (
	"fmt"
)

// 리턴값 여러 개 사용
// 함수에서 값을 여러 개 리턴할 수 있음
// ()(괄호) 안에 리턴값 자료형을 ,(콤마)로 구분하여 여러 개 지정해 줌
// 함수 안에서 값을 리턴할 때도 return에 값을 ,(콤마)로 구분하여 여러 개 지정
func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff)
}
