// 7-1
package main

import (
	"fmt"
)

// 고루틴
// 고루틴은 함수를 동시에 실행시키는 기능임
// 다른 언어의 스레드 생성보다 문법이 간단하고, 스레드보다 운영체제의 리소스를 적게 사용하므로 많은 수의 고루틴을 쉽게 생성할 수 있음

func hello() {
	fmt.Println("Hello World!")
}

func main() {
	go hello() // 함수를 고루틴으로 실행

	fmt.Scanln() // main() 함수가 종료되지 않도록 대기
}
