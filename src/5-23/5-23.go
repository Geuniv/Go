// 5-23
package main

import (
	"fmt"
)

func f() {
	a := [...]int{1, 2}

	for i := 0; i < 5; i++ { // 배열 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello World!") // 런타임 에러로 실행되지 않음
}
