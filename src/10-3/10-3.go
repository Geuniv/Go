// 10-3
package main

import (
	"fmt"
	"log"
)

// 패닉이 발생하더라도 recover 함수로 복구를 사용하면 프로그램을 종료하지 않고  넘어갈 수 있음

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}
	return "", fmt.Errorf("%d는 1이 아닙니다.", n)
}

func main() {
	defer func() { // 런타임 에러가 발생하면 recover 함수가 실행됨
		s := recover() // log.Panic 함수에서 출력한 에러 문자열 리턴
		fmt.Println(s)
	}()

	s, err := helloOne(1) // 매개변수에 1을 넣으므로 정상 동작
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(s) // Hello 1

	s, err = helloOne(2) // 매개변수에 2을 넣으므로 에러 발생
	if err != nil {
		log.Panic(err) // 에러 문자열이 출력되고 런타임 에러 발생
	}

	// 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)
	fmt.Println("Hello World !")
}
