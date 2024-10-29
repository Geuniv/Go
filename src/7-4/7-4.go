// 7-4
package main

import (
	"fmt"
	"runtime"
)

// 클로저를 고루틴으로 실행
// 예제의 출력 결과를 확실하게 표현하기 위해 CPU 코어는 하나만 사용하도록 설정

func main() {
	runtime.GOMAXPROCS(1) // CPU를 하나만 사용

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) { // 익명의 고루틴으로 실행(클로저)
			fmt.Println(s, n) // s와 매개변수로 받은 n 값 출력
		}(i) // 반복문의 변수는 매개변수로 넘겨줌
	}
	fmt.Scanln()
}
