// 7-22
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { // 데이터 구조체 정의
	tag    string // 풀 태그
	buffer []int  // 데이터 저장용 슬라이스
}

// 풀 사용
// 풀은 객체를 사용한 후 보관해두었다가 다시 사용하게 해 주는 기능임
// 객체를 반복해서 할당하면 메모리 사용량도 늘어나고, 메모리를 해제해야 하는           가비지 컬렉션에게도 부담이 됨
// 풀은 일종의 캐시라고 할 수 있으며 메모리 할당과 해제 횟수를 줄여 성능을 높이고자 할 때 사용함
// 풀은 여러 고루틴에 동시에 사용할 수 있음

// 풀은 sync.Pool을 할당한 뒤에 Get, Put 함수를 사용함
// Sync.Pool을 할당한 뒤 New 필드에 초기화 함수를 만들어 줌
// New 필드에 정의된 함수는 Get 함수를 사용했을 때 호출됨. 단 풀에 객체가 없을 때만 호출되므로 객체를 생성하고 메모리를 할당하는 코드를 작성함
// New 필드의 함수에서 new(Data)로 메모리를 할당했으므로 포인터 형인 (*Data)로 변환함
// 객체 사용이 끝났다는 의미에서 tag 필드에 used를 대입하고 Put 함수를 사용하여 객체를 풀에 보관함

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{ // 풀 할당
		New: func() interface{} { // Get 함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	for i := 0; i < 10; i++ {
		go func() { // 고루틴 10개 생성
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴

			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100) // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체를 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() { // 고루틴 10개 생성
			data := pool.Get().(*Data) // 풀에서 *Data 타입으로 데이터를 가져옴

			n := 0
			for index := range data.buffer {
				data.buffer[index] = n // 슬라이스에 랜덤 값 저장
				n += 2
			}
			fmt.Println(data)
			data.tag = "used" // 객체가 사용되었다는 태그 설정
			pool.Put(data)    // 풀에 객체를 보관
		}()
	}
	fmt.Scanln()
}
