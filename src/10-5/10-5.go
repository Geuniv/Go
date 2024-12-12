// 10-5
package main

import (
	"fmt"
	"log"
	"time"
)

// 에러 타입
// - 에러를 확장하여 더 자세한 에러 내용을 저장하기 위해 사용함

type HelloOneError struct {
	time  time.Time // 시간
	value int       // 에러를 발생시킨 값
}

func (e HelloOneError) Error() string { // 에러 메시지를 생성하여 리턴하는 Error 함수 구현
	return fmt.Sprintf("%v: %d는 1이 아닙니다.", e.time, e.value)
}

func helloOne(n int) (string, error) {
	if n == 1 { // 1일 때만 Hello 문자열을 리턴
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}
	return "", HelloOneError{time.Now(), n} // 1이 아닐 때는 에러 구조체를 생성하여 리턴
}

func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣으므로 정상 동작
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s) // Hello 1

	s, err = helloOne(2) // 매개변수에 2을 넣으므로 에러 발생
	if err != nil {
		log.Fatal(err) // 에러 문자열이 출력되고 런타임 에러 발생
	}

	// 런타임 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
	fmt.Println(s)

	fmt.Println("Hello World !")
}
