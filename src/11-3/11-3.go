// 11-3
package main

/*
#include <stdlib.h>

typedef struct _PERSON {
	char *name; // 문자열 포인터
	int age;
} PERSON;

PERSON* create(char *name, int age) // 메모리 할당 함수
{
	PERSON *p = (PERSON *)malloc(sizeof(PERSON)); // PERSON 크기로 메모리 할당

	p->name = name; // 값 설정
	p->age = age; // 값 설정
	return p;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// C 언어의 구조체 및 문자열 포인터 사용
// - C 언어의 구조체 및 문자열 포인터에 메모리를 할당한 뒤 해제함

func main() {
	var p *C.PERSON

	name := C.CString("Maria")
	age := C.int(20)

	p = C.create(name, age)

	fmt.Println(C.GoString(p.name)) // char *는 C.GoString 함수로 변환하여 사용
	fmt.Println(p.age)

	C.free(unsafe.Pointer(name)) // C.String으로 만든 문자열은 반드시 해제
	C.free(unsafe.Pointer(p))    // C 언어의 malloc 함수로 할당한 메모리는 반드시 해제
}
