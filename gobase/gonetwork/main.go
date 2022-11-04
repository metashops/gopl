package main

import (
	`fmt`
	`net`
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var boby [100]byte
	for true {
		_, err := conn.Read(boby[:])
		if err != nil {
			break
		}
		fmt.Println("receive messages: %s\n", string(boby[:]))
		_, err = conn.Write(boby[:])
		if err != nil {
			break
		}
	}
}
func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}

	// // 2、阻塞直到，拿到新的 socket
	// conn, err := ln.Accept()
	// if err != nil {
	// 	panic(err)
	// }
	// // 3、拿到新的连接就可以在业务里进行执行
	// go handleConnection(conn)

	// 3.1 需要实时监听
	for true {
		// 2、阻塞直到，拿到新的 socket
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		// 3、拿到新的连接就可以在业务里进行执行
		go handleConnection(conn)
	}
}
