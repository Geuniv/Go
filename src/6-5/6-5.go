// 6-5
package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func main() {
	var rect1 Rectangle
	var rect2 *Rectangle = new(Rectangle)

	rect1.height = 20 // 구조체 인스터스 필드에 접글할 때 .을 사용
	rect2.height = 62 // 구조체 포인터에 접근할 때도 .을 사용

	fmt.Println(rect1)
	fmt.Println(rect2) // 구조체 포인터이므로 앞에 &가 붙음
}
