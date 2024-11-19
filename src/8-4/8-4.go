// 8-4
package main

import (
	"fmt"
	"os"
)

// 디렉터리의 이름은 패키지 이름과 동일하게 만듬
// calc라는 패키지가 있으면 디렉터리는 C:/go_study/src/calc가 됨
// .go 소스 파일의 이름은 패키지 이름과 같지 않아도 됨
// 소스 파일의 첫 줄에는 package calc로 설정하여 현재 파일이 calc 패키지에 포함된 것을 알려줌
// 패키지 안에서 함수, 변수, 상수의 이름을 정하는 방법은 두 가지가 있음
// 첫 글자를 영문 소문자로 지정하면 패키지 안에서만 사용할 수 있음
// 첫 글자를 영문 대문자로 지정하면 외부에서 사용할 수 있음

// import로 calc 패키지를 가져온 뒤 calc.Sum과 같이 패키지의 함수를 사용하면 됨

// 파일 쓰기
// hello.txt 파일에 Hello, world! 문자열을 저장함

func main() {
	file, err := os.Create("hello.txt") // hello.txt 파일 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello, world !"

	n, err := file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")
}
