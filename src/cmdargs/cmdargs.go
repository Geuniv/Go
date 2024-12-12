// cmdargs
package main

import (
	"fmt"
	"os"
)

// 명령줄 옵션
// 프로그램을 실행할 때 옵션을 설정하고 싶을 때, Go 언어에서는 명령줄에서 설정한 옵션을 간단하게 사용할 수 있음
// - Go 언어 main 함수에는 매개변수가 없어서 명령줄에서 설정한 옵션을 가져오려면 os.Args 슬라이스를 사용함
// - 실행 파일을 실행하면 os.Args의 첫 번째 요소는 실행파일이며 두 번째 요소부터 옵션임

func main() {
	fmt.Println(os.Args)
}
