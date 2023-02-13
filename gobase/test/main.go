package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var body [100]byte
	for true {
		// 接收用户发来的消息
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("received: %s\n", string(body[:]))

		// 不断写回去
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
	}

}
func main() {
	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	for true {
		// 拿到新的 socket，专门为客户通信的
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
