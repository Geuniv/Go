// 5-19
package main

import (
	"fmt"
	"os"
)

func ReadHello() {
	file, err := os.Open("Hello.txt")
	defer file.Close() // 지연 호출한 file.Close()가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	ReadHello()
}
