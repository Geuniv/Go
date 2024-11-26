// 8-17
package main

import (
	"encoding/json"
	"fmt"
)

// 데이터를 JSON 형태로 변환함

// JSON 등으로 외부에 노출할 때는 구조체 필드의 첫 글자는 대문자로 만듦
// 구조체 필드의 첫 글자를 소문자로 만들면 JSON 문자에 해당 필드는 포함되지 않음

type Author struct {
	Name  string
	Email string
}
type Comment struct {
	Id      uint64
	Author  Author
	Content string
}
type Article struct {
	Id         uint64
	Title      string
	Author     Author
	Content    string
	Recommends []string
	Comments   []Comment
}

func main() {
	data := make([]Article, 1)

	data[0].Id = 1
	data[0].Title = "Hello, world !"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.MarshalIndent(data, "", "   ")

	fmt.Println(string(doc))
}
