// 8-13
package main

import (
	"bufio"
	"fmt"
	"os"
)

// 읽기, 쓰기 인터페이스를 함께 사용
// Io.ReadWriter 인터페이스로 읽기/쓰기를 처리함
// Io.ReadWriter의 정의는 다음과 같음

// Io.ReadWriter 인터페이스를 사용하여 파일에 문자열을 쓴 뒤 다시 읽음

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

	r := bufio.NewReader(file) // file로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성

	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	rw := bufio.NewReadWriter(r, w) // r, w를 사용하여 io.ReadWriter 인터페이스를 따르는 읽기/쓰기 인스턴스 생성

	rw.WriteString("Hello, world !") // 읽기/쓰기 인스턴스로 버퍼에 Hello, world ! 쓰기
	rw.Flush()                       // 버퍼의 내용을 파일에 저장

	file.Seek(0, os.SEEK_SET)   // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	data, _, _ := rw.ReadLine() // 파일에서 문자열 한 줄을 읽어서 data에 저장
	fmt.Println(string(data))
}
