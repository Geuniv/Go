// 8-5
package main

import (
	"fmt"
	"os"
)

// 파일 읽기

func main() {
	file, err := os.Open("../8-4/hello.txt") // hello.txt 파일을 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	fi, err := file.Stat() // 파일 정보를 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	n, err := file.Read(data) // 파일의 내용을 읽어서 바이스 슬라이스에 저장함
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용을 출력
}
