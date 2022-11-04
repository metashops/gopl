package main

import (
	`net/http`
)

func main() {
	server := http.Server{
		Addr:    "0.0.0.0:9090",
		Handler: nil,
	}
	server.ListenAndServe()
	// HTTPS 提供服务
	// cert.pem 是SSL证书，key.pem是服务器的私钥
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
