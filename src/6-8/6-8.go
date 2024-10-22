// 6-8
package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func (rect *Rectangle) area() int { // 리시버 변수 정의
	return rect.width * rect.height // 리시버 변수로 인스턴스에 접근
}

func main() {
	rect := Rectangle{10, 20}
	fmt.Println(rect.area())
}
