// 8-2
package main

import (
	"fmt"
	"reflect"
)

// 리플렉션으로 변수의 타입뿐만 아니라 값에 대한 상세한 정보도 얻을 수 있음

func main() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)  // f의 타입 정보를 t에 저장
	v := reflect.ValueOf(f) // f의 값 정보를 v에 저장

	fmt.Println(t.Name())                    // 자료형 이름 출력
	fmt.Println(t.Size())                    // 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // 자료의 종류를 알아 reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64)   // 자료의 종류를 알아 reflect.Int64와 비교

	fmt.Println(v.Type())                    // 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // 값 자료의 종류를 알아 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64)   // 값 자료의 종류를 알아 reflect.Int64와 비교
	fmt.Println(v.Float())                   // 값을 실수형으로 출력
}
