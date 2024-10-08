// 5-24
package main

import (
	"fmt"
)

func f() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()

	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello World!") // 복구되어서 정상적으로 실행됨
}
