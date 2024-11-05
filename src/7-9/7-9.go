// 7-9
package main

import (
	"fmt"
)

// range와 close 사용
// for 반복문 안에서 range 키워드를 사용하면 채널이 닫힐 때까지 반복하면서 값을 꺼냄

func main() {
	c := make(chan int) // int형 채널을 생성

	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄

		}
		close(c) // 채널을 닫음
	}()

	for i := range c { // range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i)
	}
}
