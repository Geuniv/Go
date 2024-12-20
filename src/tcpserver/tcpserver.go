// tcpserver
package main

import (
	"fmt"
	"net"
)

// Go 언어는 다양한 네트워크 프로토콜을 제공함
// - 그 중 TCP 프로토콜은 실시간 처리가 중요한 게임 서버에서 많이 사용함
// 서버 작성
// - 클라이언트에서 보낸 메시지를 다시 응답해주는 TCP 서버를 만듦

func requestHandler(c net.Conn) {
	data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성
	for {
		n, err := c.Read(data) // 클라이언트에서 받은 데이터를 읽음
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n])) // 데이터를 출력
		_, err = c.Write(data[:n])    // 클라이언트로 데이터를 보냄
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000") // TCP 프로토콜에 8000 포트로 연결을 받음
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close() // main 함수가 끝나기 직전에 연결 대기를 받음
	for {
		conn, err := ln.Accept() // 클라이언트가 연결되면 TCP 연결을 리턴
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer conn.Close()      // main 함수가 끝나기 직전에 TCP 연결을 닫음
	go requestHandler(conn) // 패킷을 처리할 함수를 고루틴으로 실행
}
