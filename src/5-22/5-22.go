// 5-22
package main

import (
	"fmt"
)

func f() {
	defer func() {
		s := recover() // 패닉이 발생해도 프로그램을 종료하지 않음
		// panic 함수에서 설정한 에러 메시지를 받아옴
		fmt.Println(s)
	}()
	panic("Error !!!") // panic 함수로 에러 메시지 설정, 패닉 발생

}

func main() {
	f()
	fmt.Println("Hello World!")
}
