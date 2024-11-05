// 7-17
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 고루틴 두 개에서 각각 1,000번씩 슬라이스에 값을 추가함
// data 슬라이스에 1을 2,000번 추가 했으므로 길이가 2,000이 되어야하는 데 그렇지 않음
// 두 고루틴이 경합을 벌이면서 동시에 data에 접근했기 때문에 append 함수가 정확하게 처리되지 않는 상황임. 이러한 상황을 경쟁 조건이라고 함

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // 고루틴에서 data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴에서 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // 고루틴에서 data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴에서 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기
	fmt.Println(len(data))      // data 슬라이스의 길이 출력
}
