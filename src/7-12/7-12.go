// 7-12
package main

import (
	"fmt"
)

// 채널을 리턴하려면 먼저 make 함수로 채널을 생성함. 그리고 고루틴 안에서 채널에 값을 보낸 뒤 고루틴 바깥에서 채널을 리턴함
// 채널을 리턴값으로 받았으면 <-c처럼 값을 꺼내면 됨

// 채널만 사용하여 값을 더함

func num(a, b int) <-chan int { // 함수의 리턴값은 int형 받기 전용 채널
	out := make(chan int) // int 형 채널 생성
	go func() {
		out <- a   // 채널에 a의 값을 보냄
		out <- b   // 채널에 b의 값을 보냄
		close(out) // 채널을 닫음
	}()
	return out
}

func sum(c <-chan int) <-chan int { // 채널의 매개변수는 int형 받기 전용 채널
	out := make(chan int) // int 형 채널 생성
	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼냄
			r = r + i
		}
		out <- r // 더한 값을 채널에 보냄
	}()
	return out // 채널 변수 자체를 리턴
}

func main() {
	c := num(1, 2) // 1과 2가 들어 있는 채널이 리턴됨
	out := sum(c)  // 채널 c를 매개변수에 넘겨서 모두 더함

	fmt.Println(<-out) // out 채널에서 값을 꺼냄
}
