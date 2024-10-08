// 5-3
package main

import (
	"fmt"
)

// 리턴값에 이름을 지정할 수 있음
// 괄호 안에 리턴값 변수의 이름과 자료형을 지정함
// 리턴값 변수를 사용할 때는 return 뒤에 리턴할 변수를 지정하지 않음
func sum(a int, b int) (r int) {
	r = a + b
	return
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
