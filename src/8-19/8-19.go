// 8-19
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSON 파일 사용
// - 데이터를 JSON 파일로 저장

type Author struct {
	Name  string `json:"name"` // 구조체 필드에 태그 지정
	Email string `json:"email"`
}

type Comment struct {
	Id      uint64 `json:"id"`
	Author  Author `json:"author"`
	Content string `json:"content"`
}

type Article struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     Author    `json:"author"`
	Content    string    `json:"content"`
	Recommends []string  `json:"recommends"`
	Comments   []Comment `json:"comments"`
}

func main() {
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

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

	doc, _ := json.MarshalIndent(data, "", "   ") // 맵을 JSON 문서로 변환
	// 사람이 쉽게 읽을 수 있도록 변환

	err := ioutil.WriteFile("./article.json", doc, os.FileMode(0644))
	// article.json 파일에 JSON 문서 저장

	if err != nil {
		fmt.Println(err)
		return
	}
}
