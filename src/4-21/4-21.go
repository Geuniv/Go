// 4-21
package main

import (
	"fmt"
)

func main() {
	a := map[string]int{"Hello": 10, "world": 20}

	b := map[string]int{
		"Hello": 10,
		"world": 20, // 여러 줄로 나열할 때 마지막 요소에 콤마를 붙임
	}
	fmt.Println(a["Hello"])
	fmt.Println(b["world"])
}
