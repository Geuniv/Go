// 7-7
package main

import (
	"fmt"
	"time"
)

// 동기 채널
// 아래의 예제는 고루틴과 메인 함수를 번갈아 가면서 실행함

func main() {
	done := make(chan bool) // 동기 채널 생성
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 고루틴에 true를 보냄. 값을 꺼낼 때까지 대기
			fmt.Println("고루틴 : ", i)    // 번복문의 변수 출력
			time.Sleep(1 * time.Second) // 1초 대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done                    // 채널에 값이 들어올 때까지 대기, 값을 꺼냄
		fmt.Println("메인함수 : ", i) // 반복문의 변수 출력
	}
}
