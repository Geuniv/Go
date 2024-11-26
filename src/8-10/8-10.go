// 8-10
package main

import (
	"io"
	"os"
	"strings"
)

// 문자열을 화면에 출력
// 문자열 io.Reader를 그대로 화면에 출력함

func main() {
	s := "Hello, world !"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성

	io.Copy(os.Stdout, r) // os.Stdout에 io.Reader를 복사하면 화면에 그대로 출력됨
}
