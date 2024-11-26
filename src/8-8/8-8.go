// 8-8
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Go 언어는 io.Reader, io.Writer 인터페이스를 활용하여 다양한 방법으로 입출력을 할 수 있음

// io.Reader와 io.Writer의 정의는 다음과 같음
// 어떤 구조체든 매개변수로 바이트 슬라이스를 받고, 정수와 에러 값을 리턴하는 Read 함수를 가지고 있으며 io.Reader 인터페이스를 따른다고 할 수 있음
// 매개변수로 바이트 슬라이스를 받고, 정수와 에러 값을 리턴하는 Write 함수를 가지고 있으며 io.Writer 인터페이스를 따른다고 할 수 있음

// 파일 처리
// bufio는 Buffered I/O를 의미하며 io.Reader, io.Writer 인터페이스를 받음

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,

		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	w.WriteString("Hello, world !") // 쓰기 인스턴스로 버퍼에 Hello, world 쓰기
	w.Flush()                       // 버퍼의 내용을 파일에 저장

	r := bufio.NewReader(file) // io.Reader 인터페이스를 따르는 인스턴스 r 생성

	fi, _ := file.Stat()         // 파일의 정보 구함
	b := make([]byte, fi.Size()) // 파일의 크기만큼 슬라이스 생성

	file.Seek(0, os.SEEK_SET) // 파일 읽기 위치를 파일의 맨 처음 ( 0 ) 으로 이동
	r.Read(b)                 // 읽기 인스턴스로 파일의 내용을 읽어서 b에 저장

	fmt.Println(string(b)) // 문자열로 변환하여 b의 내용을 저장
}
