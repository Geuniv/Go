// 7-13
package main

import (
	"fmt"
	"time"
)

// 셀렉트 사용
// 여러 채널을 손쉽게 사용할 수 있도록 select 분기문을 제공함
// 보통 select를 계속 처리할 수 있도록 for로 반복해 줌

// 채널 2개를 생성하고 100밀리초, 500밀리초 간격으로 숫자와 문자열을 보낸 뒤 꺼내서 출력함

func main() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string형 채널 생성
	go func() {
		for {
			c1 <- 10                           // 채널 c1에 10을 보낸 뒤
			time.Sleep(100 * time.Millisecond) // 100밀리초 대기
		}
	}()
	go func() {
		for {
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!을 보낸 뒤
			time.Sleep(500 * time.Millisecond) // 500밀리초 대기
		}
	}()
	go func() {
		for {
			select {
			case i := <-c1: // 채널 c1에 값이 들어오면 값을 꺼내 i에 대입
				fmt.Println("c1 : ", i)
			case s := <-c2: // 채널 c2에 값이 들어오면 값을 꺼내 s에 대입
				fmt.Println("c2 : ", s)
			}
		}
	}()
	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
