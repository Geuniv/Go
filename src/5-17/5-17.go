// 5-17
package main

import (
	"fmt"
)

func HelloWorld() {
	defer func() { // HelloWorld() 함수가 끝나기 직전에 호출
		fmt.Println("world!")
	}()
	func() {
		fmt.Println("Hello")
	}()
}

func main() {
	HelloWorld()
}
