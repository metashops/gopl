package tcp

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go-redis/interface/tcp"
	"go-redis/lib/logger"
)

type Config struct {
	Address string
}

func ListenAndServerWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	// cfg.Address = listener.Addr().String()
	logger.Info(fmt.Sprintf("bind: %s, start listening...", cfg.Address))
	ListenAndServer(listener, handler, closeChan)
	return nil
}

func ListenAndServer(
	listener net.Listener,
	handler tcp.Handler,
	closeChan <-chan struct{}) {

	go func() {
		<-closeChan
		logger.Info("shutting down")
		_ = listener.Close()
		_ = handler.Close()
	}()

	defer func() {
		_ = listener.Close()
		_ = handler.Close()
	}()

	ctx := context.Background()
	var waitDone sync.WaitGroup
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accept link")
		waitDone.Add(1)
		go func() {
			defer func() {
				waitDone.Done()
			}()
			handler.Handler(ctx, conn)
		}()
	}
	waitDone.Wait()
}
