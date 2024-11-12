// 7-24
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 원자적 연산 사용
// 원자적 연산은 더 이상 쪼갤 수 없는 연산을 의미함
// 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있음

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1) // 원자적 연산으로 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1) // 원자적 연산으로 1씩 뺌
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(data)
}
