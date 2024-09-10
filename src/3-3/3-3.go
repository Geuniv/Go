// 3-3
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		str1 := string(b[:])
		fmt.Println(str1)
	}
}
