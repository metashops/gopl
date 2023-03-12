package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1.
	r.GET("/slow", SlowHandler)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go server.ListenAndServe()
	// 设置优雅退出
	gracefulExitWeb(server)
}

func SlowHandler(c *gin.Context) {
	fmt.Println("[start] SlowHandler")
	// time.Sleep(30 * time.Second)
	time.Sleep(10 * time.Second)
	fmt.Println("[end] SlowHandler")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil {
		fmt.Println("err", err)
	}

	// 看看实际退出所耗费的时间
	fmt.Println("------exited--------", time.Since(now))
}
