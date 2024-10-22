// 6-9
package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello ~")
}

type Student struct {
	p      Person // 학생 구조체 안에서 사람 구조체를 필드로 가지고 있음
	school string
	grade  int
}

func main() {
	var s Student

	s.p.greeting()
}
