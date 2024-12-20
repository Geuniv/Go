// 8-20
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// JSON 문서를 저장한 article.json 파일을 읽음

type Author struct {
	Name  string `json:"name"`
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
	b, err := ioutil.ReadFile("../8-19/article.json") // article.json 파일을 읽어서
	// 바이트 슬라이스에 저장

	if err != nil {
		fmt.Println(err)
		return
	}

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal(b, &data) // JSON 문서의 내용을 변환하여 data에 저장
	fmt.Println(data)
}
