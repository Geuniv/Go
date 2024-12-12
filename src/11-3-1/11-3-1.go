// 11-3-1
package main

// C 언어에서 Go 언어의 함수를 사용함

// C 언어에서 매개변수를 넣고 리턴값을 사용할 수록 매개변수와 리턴값의 타입을   각각 C.int로 지정함
// C.int는 C 언어의 int 자료형이며 Go 언어의 int와는 다른 자료형임
// 함수 정의 윗 부분에 //export 함수명 형식으로 주석을 작성함
// - 이 주석이 있어야 C 언어에서 Go 언어의 함수를 사용할 수 있음
// - // 뒤에는 띄어 쓰지 않고 붙여줌
// - 여기서는 sum함수 이므로 export sum임
// extern int sum(int a, int b);으로 외부에 sum함수가 있다는 것을 알려줌. 여기서  외부는 Go 언어임
// Go 언어에서 함수를 //export 로 설정한 뒤 주석 형식으로 C 언어 함수를 작성할 때는 반드시 static inline으로 지정해줌

/*
#include <stdio.h>

extern int sum(int a, int b); // Go 언어의 함수는 extern으로 선언

static inline void CExample() {
	int r = sum(1, 2); // Go 언어의 sum 함수 호출
	printf("%d\n", r);
}
*/
import "C"

//export sum
func sum(a, b C.int) C.int { // C 언어에서 사용할 수 있도록 매개변수와 리턴값 자료형을
	// C 언어용으로 맞춤
	return a + b
}

func main() {
	C.CExample()
}
