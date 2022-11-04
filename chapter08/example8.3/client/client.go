package main

import (
	`io`
	`log`
	`net`
	`os`
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdout)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
