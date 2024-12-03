// 9-1
package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

// 9장) 압축, 암호화, 정렬과 자료구조
// Go 언어는 압축 알고리즘을 패키지로 제공함
// - gzip 알고리즘을 사용해서 데이터를 압축한 뒤 파일로 저장함

func main() {
	file, err := os.OpenFile( // 압축할 파일 생성
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	w := gzip.NewWriter(file) // file로 io.Writer 인터페이스를 따르는
	// 압축 인스턴스 w 생성

	defer w.Close() // main 함수가 끝나기 직전에 압축 인스턴스를 닫음

	s := "Hello, world !"
	w.Write([]byte(s)) // 압축 인스턴스로 문자열을 압축하여 파일에 저장
	w.Flush()          // 메모리 상에 데이터를 파일에 완전히 저장
}
