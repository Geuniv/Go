// 6-12
package main

import (
	"fmt"
)

// 인터페이스는 메서드의 집함임
// 인터페이스는 메서드 자체를 구현하지 않음
type hello interface {
}

func main() {
	var h hello
	fmt.Println(h)
}
