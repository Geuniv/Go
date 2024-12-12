// 10-2
package main

import (
	"fmt"
	"log"
)

// 에러 처리
// fmt.Errorf, log.Factal 함수를 사용하여 에러를 처리하는 방법
// - 숫자가 1일 때 문자열을 리턴하는 함수가 있는데 1이 아닐 때는 의도적으로 에러를 발생시킴
// - 프로그램을 정상 종료하지 않고 런타임 에러를 발생시키려면 log.Panic 함수를 사용

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}
	return "", fmt.Errorf("%d는 1이 아닙니다.", n)
}

func main() {
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
