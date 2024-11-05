// 7-18
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 보호를 시작할 부분에서 Lock 함수를 사용하고, 보호를 끝낼 부분에서 Unlock 함수를 사용함
// Lock, Unlock 함수는 반드시 짝을 맞지 않으면 데드락(교착 상태)이 발생함

// 읽기, 쓰기 뮤텍스 사용
// 읽기, 쓰기 뮤텍스는 읽기 동작과 쓰기 동작을 나누어 잠금(락)을 걸 수 있음
// 읽기 락 : 읽기 락끼리는 막지 않음. 하지만 읽기 시도 중에 값이 바뀌면 안되므로  쓰기 락은 막음
// 쓰기 락 : 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안되고, 다른 곳에서 값을 바꾸면 안되므로 읽기, 쓰기 락 모두 막음
// read1, read2 읽기 동작이 모두 끝나야 write 쓰기 동작이 시작되고, 쓰기 동작이  끝나야 읽기 동작이 시작됨
// 읽기 동작끼리는 서로 막지 않으므로 항상 동시에 실행됨

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	var data int = 0
	var rwMutex = new(sync.RWMutex) // 읽기, 쓰기 뮤텍스 생성

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock() // 쓰기 뮤텍스 잠금, 쓰기 보호 시작
			data += 1      // data에 값 쓰기
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond) // 10밀리초 대기
			rwMutex.Unlock()                  // 쓰기 뮤텍스 해제, 쓰기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second) // 1초 대기
			rwMutex.RUnlock()           // 읽기 뮤텍스 해제, 읽기 보호 종료
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock() // 읽기 뮤텍스 잠금, 읽기 보호 시작
			fmt.Println("read 2 : ", data)
			time.Sleep(1 * time.Second) // 2초 대기
			rwMutex.RUnlock()           // 읽기 뮤텍스 해제, 읽기 보호 종료
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
