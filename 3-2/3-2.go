// 3-2
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var b []byte
	var err error

	b, err = ioutil.ReadFile("./hello.txt")

	if err == nil {
		str1 := string(b[:])
		fmt.Println(str1)
	}
}
