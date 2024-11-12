// 7-16
package main

import (
	"fmt"
	"runtime"
	"time"
)

// 동기화 객체
// Go 언어는 채널 이외에도 고루틴의 실행 흐름을 제어하는 동기화 객체를 제공함
// 대표적인 동기화 객체
// 뮤텍스 : 상호 배제라고도 하며 여러 스레드(고루틴)에서 공유되는 데이터를 보호할 때 주로 사용
// 읽기, 쓰기 뮤텍스 : 읽기와 쓰기 동작을 나누어서 잠금(락)을 걸 수 있음
// 조건 변수 : 대기하고 있는 하나의 객체를 깨울 수 있고 여러 개를 동시에 깨울 수  있음
// 함수를 한번만 실행 : 특정 함수를 딱 한번만 실행할 때 사용함
// 풀 : 멀티스레드(고루틴)에서 사용할 수 있는 객체 풀. 자주 사용하는 개체를 풀에  보관했다가 다시 사용함
// 대기 그룹 : 고루틴이 모두 끝날 때까지 기다리는 기능
// 원자적 연산 : 더 이상 쪼갤 수 없는 없는 연산. 멀티스레드(로루틴), 멀티코어 환경에서 안전하게 값을 연산하는 기능

// 뮤텍스(Mutex) 사용
// 뮤텍스는 여러 고루틴이 공유하는 데이터를 보호할 때 사용함

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{} // int 형 슬라이스 생성

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) // 고루틴에서 data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴에서 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			data = append(data, 1) // 고루틴에서 data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴에서 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기
	fmt.Println(len(data))      // data 슬라이스의 길이 출력
}