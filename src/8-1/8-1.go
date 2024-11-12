// 8-1
package main

import (
	"fmt"
	"reflect"
)

// 리플렉션은 실행 시점에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는  기능임

type Data struct { // 구조체 정의
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // reflect.TypeOf로 자료형 이름 출력

	var s string = "Hello, world !"
	fmt.Println(reflect.TypeOf(s)) // reflect.TypeOf로 자료형 이름 출력

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f)) // reflect.TypeOf로 자료형 이름 출력

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // reflect.TypeOf로 자료형 이름 출력
}
