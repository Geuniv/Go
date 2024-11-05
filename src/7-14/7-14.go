// 7-14
package main

import (
	"fmt"
	"time"
)

// select 분기문은 채널에 값을 보낼 수도 있음

func main() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string형 채널 생성
	go func() {
		for {
			i := <-c1 // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 : ", i)
			time.Sleep(100 * time.Millisecond) // 100밀리초 대기
		}
	}()
	go func() {
		for { // 무한 루프
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!을 보낸 뒤
			time.Sleep(500 * time.Millisecond) // 500밀리초 대기
		}
	}()
	go func() {
		for {
			select {
			case c1 <- 10: // 매번 채널 c1에 10을 보냄
			case s := <-c2: // 채널 c2에 값이 들어오면 값을 꺼내 s에 대입
				fmt.Println("c2 : ", s)
			}
		}
	}()
	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
