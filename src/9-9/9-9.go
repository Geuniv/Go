// 9-9
package main

import (
	"container/list"
	"fmt"
)

// Go에서 제공하는 자료구조는 다음과 같음
// - 연결 리스트 : 각 노드를 한 줄로 연결한 자료구조임
// - 힙 : 이진 트리를 활용한 자료구조로 부모 노드에서 자식 노드가 대소 관계를 이루며 정렬이 됨
// - 링 : 각 노드가 원형으로 연결된 자료구조

// 연결 리스트
// - 연결 리스트를 생성한 뒤 데이터를 넣고, 각 노드를 순회함
// - Go 언어의 연결 리스트는 이중 연결 리스트임. 따라서 앞, 뒤 양쪽 방향으로 순회함

func main() {
	l := list.New() // 연결 리스트 생성
	l.PushBack(10)  // 연결 리스트 데이터 추가
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value) // 연결 리스트의 맨 앞 데이터를 가져옴
	fmt.Println("Back ", l.Back().Value)   // 연결 리스트의 맨 뒤 데이터를 가져옴

	for e := l.Front(); e != nil; e = e.Next() { // 연결 리스트의 맨 앞부터 끝까지 순회
		fmt.Println(e.Value)
	}
}
