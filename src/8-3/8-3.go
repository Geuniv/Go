// 8-3
package main

import (
	"fmt"
	"reflect"
)

// 리플렉션을 사용하여 함수를 만들어 내는 방법
// 리플렉션으로 생성하는 함수는 []reflect.Value를 매개변수로 받고 리턴값으로 사용해야 함
// 함수를 담을 수 있는 변수 hello를 선언하고 reflect.ValueOf 함수에 hello의 주소를 넘긴 뒤 Elem 함수로 실제 값 정보를 가져옴
// Reflect.MakeFunc(fn.Type(), h)와 같이 함수 h를 넣어서 함수 정보를 생성함
// fn.Set(v)와 같이 hello의 값 정보인 fn에 함수 정보 v를 설정하여 함수를 연결해 줌

func h(args []reflect.Value) []reflect.Value { // 매개변수와 리턴값은 []reflect.Value를 사용
	fmt.Println("Hello World !")
	return nil
}

func main() {
	var hello func() // 함수를 담은 변수 선언

	fn := reflect.ValueOf(&hello).Elem() // hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), h) // h의 함수 정보를 생성

	fn.Set(v) // hello의 값 정보인 fn에 h의 함수 정보를 v를 설정하여 함수를 연결

	hello()
}
