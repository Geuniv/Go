// 5-1
package main

import (
	"fmt"
)

// 함수 정의를 시작한 줄에서 {(여는 중괄호)가 시작됨
// 여는 중괄호를 다음 줄에 작성하면 컴파일 에러가 발생함
func hello() {
	fmt.Println("Hello World!")
}

func main() {
	hello()
}
