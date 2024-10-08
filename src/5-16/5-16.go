// 5-16
package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // 현재 함수 main()이 끝나기 직전에 호출
	hello()
	hello()
	hello()
}
