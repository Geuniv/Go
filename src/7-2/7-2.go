// 7-2
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 반복문을 사용하여 고루틴을 100개 생성함

func hello(n int) {
	r := rand.Intn(100)          // 랜덤한 숫자 생성
	time.Sleep(time.Duration(r)) // 랜덤한 시간동안 대기
	fmt.Println(n)
}

func main() {
	for i := 0; i < 100; i++ {
		go hello(i) // 고루틴 100개 생성
	}
	fmt.Scanln()
}
