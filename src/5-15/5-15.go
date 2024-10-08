// 5-15
package main

import (
	"fmt"
)

func calc() func(x int) int {
	a, b := 3, 5 // 지역 변수는 함수가 끝나면 소멸되지만
	return func(x int) int {
		return a*x + b // 클로저이므로 함수를 호출할 때마다 변수 a와 b 값을 사용
	}
}

func main() {
	f := calc() // calc 함수를 실행하여 리턴값으로 나온 클로저를 변수에 저장

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))
	fmt.Println(f(5))
}
