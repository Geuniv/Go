// 7-19
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 조건 변수 사용
// 조건 변수는 대기하고 있는 객체 하나만 깨우거나 여러 개를 동시에 깨울 때 사용함
// 조건 변수는 뮤텍스를 먼저 생성한 뒤 sync.NewCond 함수로 생성함
// 대기할 때는 고루틴 안에서 Wait 함수를 사용하며, 대기하는 고루틴을 깨울 때는 Signal 함수를 사용함
// Wait 함수 부분은 뮤텍스의 Lock, Unlock 함수로 보호해야 함

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 사용

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) { // 고루틴 3개 생성
			mutex.Lock() // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true    // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() // 뮤텍스 잠금 해제, cond.Wait() 보호 종료
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄. 고루틴 3개가 모두 실행될 때까지 기다림
	}

	for i := 0; i < 3; i++ {
		mutex.Lock() // 뮤텍스 잠금, cond.Signal() 보호 시작
		fmt.Println("signal : ", i)
		cond.Signal()  // 대기하고 있는 고루틴을 하나씩 깨움
		mutex.Unlock() // 뮤텍스 잠금 해제, cond.Signal() 보호 종료
	}

	fmt.Scanln()
}
