package main

import (
	`fmt`
	`net/http`
)

//  ServeHTTP 处理请求
// type MyHandler struct{}
//
// func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world!")
// }

// func main() {
// 	handler := MyHandler{}
// 	server := &http.Server{
// 		Addr:    "127.0.0.1:8080",
// 		Handler: &handler,
// 	}
// 	server.ListenAndServe()
// }

// 以上问题：
// 创建一个处理器并它与服务器进行绑定，以此来代替原本在使用的默认多路复用器
// 这样意味着服务器不会再通过URL匹配将请求路由至不同的处理器来，而是直接使用一个处理器来处理所有
// 因此无论浏览器输出访问什么地址，服务器返回都是同样的 Hello World!

// 我们更加希望通过 URL 请求返回不同的响应

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()

}
