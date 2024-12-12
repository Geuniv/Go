// 11-1
package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

// C 언어 연동
// Go 언어에서 C언어 함수 사용
// - Go 언어에서 C언어의 rand 함수를 사용함
// - rand 함수를 사용하려면 다음과 같이 주석 안에 #include <stdlib.h>처럼 헤더 파일을 포함함
// - 그리고 주석 바로 밑에 impot “C”를 넣어줌
// - 이 때 주석과 import 사이에 새 줄이 있거나 다른 코드가 있으면 안됨
// - C 언어 함수들은 항상 C.으로 시작함. 그러므로 rand 함수는 C.rand처럼 사용함

func main() {
	r := C.rand() // stdlib의 rand 함수 호출
	fmt.Println(r)
}
