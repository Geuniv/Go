// 7-5
package main

import (
	"fmt"
	"runtime"
)

// 클로저를 고루틴으로 실행할 때 반복문 안에서 변수 사용에 주의해야 함
// 반복문으로 증가하는 i를 클로저에서 그대로 사용하지 않고 매개변수로 넘겨주었음
// 일반 클로저는 반복문 안에서 순서대로 실행되지만 고루틴으로 실행한 클로저는 반복문이 끝난 뒤에 고루틴이 실행됨

// 변수 i를 클로저의 매개변수로 넘기지 않고 fmt.Println으로 바로 출력함

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문의 변수를 클로저에서 바로 출력
		}()
	}
	fmt.Scanln()
}
