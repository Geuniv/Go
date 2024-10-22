// 6-13
package main

import (
	"fmt"
)

// 메서드를 가지는 인터페이스를 정의함
// {}(중괄호), 블록 안에 메서드 이름, 매개변수 자료형, 리턴값 자료형을 지정하여 한 줄씩 나열함. ,(콤마)로 구분하지 않음

type MyInt int // int 형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5

	var p Printer // 인터페이스 선언

	p = i // i를 인터페이스 p에 대입

	p.Print() // 인터페이스를 통하여 MyInt의 Print 메서드 호출
}
