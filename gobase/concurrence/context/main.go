package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// 当ctx.Done()通道被关闭时，意味着取消信号已被触发
			fmt.Println("worker canceled")
			return
		default:
			// 执行某些操作
			fmt.Println("worker is working")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(ctx, &wg)
	}

	cancel()
	wg.Wait()

	time.Sleep(3 * time.Second)
}
