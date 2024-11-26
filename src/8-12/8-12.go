// 8-12
package main

import (
	"bufio"
	"fmt"
	"os"
)

// bufio와 기본 입출력 함수를 사용하여 파일에 값을 저장함

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,

		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	fmt.Fprintf(w, "%d, %f, %s", 1, 1.1, "Hello")

	w.Flush()
}
