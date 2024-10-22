// 6-15
package main

import (
	"fmt"
)

// 덕 타이핑
// 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드만 타입을 판단하는 방식

type Duck struct { // 오리 구조체 정의
}

func (d Duck) quack() { // 오리의 quack 메서드 정의
	fmt.Println("꽥~!")
}

func (d Duck) feathers() { // 오리의 feathers 메서드 정의
	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

type Person struct { // 사람의 구조체 정의
}

func (p Person) quack() { // 사람의 quack 메서드 정의
	fmt.Println("사람은 오리를 흉내냅니다. 꽥~!")
}

func (p Person) feathers() { // 사람의 featchers 메서드 정의
	fmt.Println("사람은 방에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface { // quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck
	var john Person

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feathers 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feathers 메서드 호출
}
