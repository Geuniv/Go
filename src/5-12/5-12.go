// 5-12
package main

import (
	"fmt"
)

func main() {
	func() { // 함수에 이름이 없음
		fmt.Println("Hello World!")
	}()

	func(s string) { // 익명의 함수를 정의한 후
		fmt.Println(s)
	}("Hello World!") // 바로 호출

	r := func(a int, b int) int { // 익명의 함수를 정의한 후
		return a + b
	}(1, 2) // 바로 호출하여 리턴값을 변수 r에 저장

	fmt.Println(r)
}
