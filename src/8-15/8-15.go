// 8-15
package main

import (
	"encoding/json"
	"fmt"
)

// JSON 문서를 Go 언어에서 읽으려면 encoding/json 패키지의 Unmarshal 함수를  사용함
// JSON 문서의 데이터를 저장할 공간이 필요한데 맵으로 만들 수 있음
// JSON 문서에서 키는 문자열이며 값은 문자열과 숫자를 사용하고 있음
// - 맵을 만들 때는 키는 문자열로 지정하고 값은 interface{}로 지정하여 모든 값을 넣을 수 있도록 함
// Json.Unmarshal 함수에 JSON 문자열을 []byte 형식으로 변환하여 넣고, 맵의          포인터도 넣음
// 맵에서 data[“name”]처럼 키를 지정하면 값을 가져올 수 있음

// 맵 형태의 데이터를 JSON 형태로 변환함

func main() {
	data := make(map[string]interface{}) // 문자열을 키로 하고 모든 자료형을 저장할 수 있는 맴 생성

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.MarshalIndent(data, "", "   ") // 맵을 JSON 문서로 변환
	// 사람이 쉽게 읽을 수 있도록 변환

	fmt.Println(string(doc))
}
