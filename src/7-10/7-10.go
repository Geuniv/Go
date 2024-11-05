// 7-10
package main

import (
	"fmt"
)

// 보내기 전용 및 받기 전용 채널 사용
// 보내기 전용 채널과 받기 전용 채널은 값의 흐름이 한 방향으로 고정된 채널임

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100 // 채널에 값을 보냄
	// fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c) // 채널에 값을 꺼냄
	//ㅊ <- 1 // 채널에 값을 보내면 컴파일 에러
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
