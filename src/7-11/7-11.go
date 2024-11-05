// 7-11
package main

import (
	"fmt"
)

func sum(a, b int) <-chan int { // 함수의 리턴값은 int형 받기 전용 채널
	out := make(chan int) // 채널 생성
	go func() {
		out <- a + b // 채널에 a와 b의 합을 보냄
	}()
	return out // 채널 변수 자체를 리턴
}

func main() {
	c := sum(1, 2) // 채널을 리턴값으로 받아서 c에 대입

	fmt.Println(<-c) // 채널에서 값을 꺼냄
}
