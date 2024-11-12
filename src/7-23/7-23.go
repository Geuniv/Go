// 7-23
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 대기 그룹 사용
// 대기 그룹은 고루틴이 끝날 때까지 기다릴 때 사용함

// 대기 그룹은 sync.WaitGroup을 할당한 뒤에 Add, Done, Wait 함수로 사용함
// 고루틴을 생성할 때 Add 함수로 고루틴 개수를 추가해 주고 고루틴 안에서           Done 함수를 사용하여 고루틴이 끝났다는 것을 알려줌
// 마지막으로 Wait 함수를 사용하여 모든 고루틴이 끝날 때까지 기다림
// Done 함수는 defer 함수와 함께 사용해서 지연 호출을 사용할 수 있음

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1)        // 반복할 때마다 wg.Add 함수로 1씩 추가
		go func(n int) { // 고루틴 10개 생성
			fmt.Println(n)
			wg.Done() // 고루틴이 끝났다는 것을 알려줌
		}(i)
	}

	wg.Wait() // 모든 고루틴이 끝날 때까지 기다림
	fmt.Println("the end")
}
