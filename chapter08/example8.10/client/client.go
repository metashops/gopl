package main

import (
	`log`
	`net`
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
