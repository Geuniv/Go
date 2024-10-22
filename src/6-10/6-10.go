// 6-10
package main

import (
	"fmt"
)

type Student struct {
	Person // 필드명 없이 선언하면 포함관계가 됨
	school string
	grade  int
}

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello ~")
}

func main() {
	var s Student

	s.Person.greeting()
	s.greeting()
}
