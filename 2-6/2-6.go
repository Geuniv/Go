// 2-6
package main

import (
	"fmt"
)

func main() {
	var s1 string = "Hello, world!\n"
	var s2 string = "안녕하세요\n"
	var s3 string = `안녕하세요
Hello, world`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
