package main

// C 언어의 배열을 Go 언어의 슬라이스로 만듦
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ARRAY_LENGTH 5

int *GetArray()
{
	int *a = (int *)malloc(sizeof(int) * ARRAY_LENGTH); // 메모리 할당
	memset(a, 0, sizeof(int) * ARRAY_LENGTH); // 0으로 초기화

	a[0] = 21; // 배열의 요소에 값을 넣어줌
	a[1] = -15;
	a[2] = 68;
	a[3] = 72;
	a[4] = -33;

	return a; // 할당한 메모리 리턴
}

int GetLength() // 배열의 길이를 구하는 함수
{
	return ARRAY_LENGTH;
}
*/
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a *C.int = C.GetArray()  // 배열의 포인터를 받음
	length := int(C.GetLength()) // 길이를 구함

	hdr := reflect.SliceHeader{ // 배열을 감싸는 슬라이스 헤더 생성
		Data: uintptr(unsafe.Pointer(a)), // unsafe.Pointer로 변환하여 넣어줌
		Len:  length,                     // 배열의 길이
		Cap:  length,
	}
	// 슬라이스 헤더를 통해 []C.int 생성
	s := *(*[]C.int)(unsafe.Pointer(&hdr)) // 슬라이스 헤더의 포인터로 변환한 뒤 역참조
	fmt.Println(s)
	fmt.Println(s[2:5])

	C.free(unsafe.Pointer(a)) // C 언어에서 malloc 함수로 할당한 메모리는 반드시 해제
}
