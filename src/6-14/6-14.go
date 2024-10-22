// 6-14
package main

import (
	"fmt"
)

// int 자료형과 사각형 구조체의 내용을 출력하고 int 자료형과 사각형 구조체의 인스턴스를 담을 수 있는 인터페이스를 정의함
// 인터페이스는 자료형이든 구조체이든 타입에 상관없이 메서드 집합만 같으면 동일한 타입으로 봄

type MyInt int // int형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(r.width, r.height)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}
	var p Printer

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = r     // range를 인터페이스 p에 대입
	p.Print() // 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출
}
