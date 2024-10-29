// 7-8
package main

import (
	"fmt"
	"runtime"
)

// 채널 버퍼링
// 채널의 버퍼가 가득차면 값을 꺼내서 출력함
// 채널에 버퍼를 1개 이상 설정하면 비동기 채널이 생성됨
// 비동기 채널은 보내는 쪽에서 버퍼가 가득 차면 실행을 멈추고 대기하며 받는 쪽에서는 버퍼에 값이 없으면 대기함

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 3) // 버퍼가 3개인 비동기 채널 생성
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true             // 채널에 true를 보냄, 버퍼가 가득 차면 대기
			fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
		}
	}()

	for i := 0; i < count; i++ {
		<-done                    // 버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인함수 : ", i) // 반복문의 변수 출력
	}
}
