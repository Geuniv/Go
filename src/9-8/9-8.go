// 9-8
package main

import (
	"fmt"
	"sort"
)

// 정수, 실수, 문자열과 같은 기본 자료형은 슬라이스에 바로 넣을 수 있는 함수를 제공
// - 기본적으로 오름차순 정렬을 함

func main() {
	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}

	sort.Ints(a)
	fmt.Println(a)

	sort.Float64s(b)
	fmt.Println(b)

	sort.Strings(c)
	fmt.Println(c)
}
