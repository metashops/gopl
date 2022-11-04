package main

import (
	`fmt`
	`net/http`
)

func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!（你好！）")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!（世界!）")
}
