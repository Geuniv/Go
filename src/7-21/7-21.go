// 7-21
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 함수를 한 번만 실행
// Once를 사용하면 함수를 한 번만 실행할 수 있음
// Once는 sync.Once를 할당한 뒤에 Do 함수로 사용함
// Do 함수에는 실행할 함수 이름을 지정하거나, 클로저 함수를 지정할 수 있음
// Once는 어떤 상황이든 상관없이 지정된 함수를 딱 한 번만 실행시킴

func hello() {
	fmt.Println("Hello World !")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			fmt.Println("goroutine : ", n)

			once.Do(hello) // 고루틴은 3개지만 hello 함수를 한 번만 실행
		}(i)
	}
	fmt.Scanln()
}
