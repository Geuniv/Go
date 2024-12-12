// swap
package main

// C 파일 사용
// 주석 형태로 C 언어 코드를 작성하지 않고 .c 파일을 따로 빼내는 방법을 알아봄
// Go 언어에서 .c 파일에 들어 있는 Cexample 함수를 호출하기 위해 extern으로 선언해 줌
// .c에서 Go 언어의 swap 함수를 호출하기 위해 extern으로 선언해 줌
// C 언어에는 리턴값을 한 개만 사용할 수 있으므로 Go 언어에서 리턴값을 두 개이상 리턴할 때는 C 구조체로 정의함
// - 리턴값 구조체는 struct 함수명_return이 됨

// C 언어 파일에는 반드시 _cgo_export.h 헤더 파일을 포함함
// - 이 파일은 컴파일 하는 동안 임시로 생성되며 현재 디렉터리에는 생성되지 않음
// struct swap_return s와 같이 구조체로 리턴값을 받을 변수를 선언한 뒤  Go 언어의 swap 함수를 호출함

// go build swap처럼 디렉터리 이름을 지정하여 모든 소스 파일(.go, .c)이 컴파일되도록 실행함
// 컴파일된 실행 파일을 실행하면 swap 함수의 리턴값이 출력됨

/*
extern void CExample(); // Go 언어에서 C 언어의 함수를 호출하기 위한 선언
extern struct swap_return swap(int a, int b);
// C 언어에서 Go 언어의 함수를 호출하기 위한 선언.
// 구조체를 이용하여 두 개의 값을 리턴
*/
import "C"

//export swap
func swap(a, b, C.int) (C.int, C.int) { // C 언어에서 사용할 함수, 리턴값이 두 개
	return b, a
}

func main() {
	C.CExample()
}
