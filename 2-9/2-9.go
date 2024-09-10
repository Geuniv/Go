// 2-9
package main

import (
	"fmt"
)

func main() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"

	fmt.Println(s1 == s2)
	fmt.Println(s1 + s2 + s3)
	fmt.Println("안녕하세요" + s1)
}
