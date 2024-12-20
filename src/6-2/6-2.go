// 6-2
package main

import (
	"fmt"
)

func main() {
	var num int = 1
	var numPtr *int = &num // 참조로 num 변수에 메모리 주소를 구하여
	// numPtr 포인터 변수에 대입

	fmt.Println(numPtr) // 0xc0000a1e8: numPtr 포인터 변수에 저장된 메모리 주소
	fmt.Println(&num)   // 0xc0000a1e8: 참조로 num 변수의 메모리 주소를 구함
}
