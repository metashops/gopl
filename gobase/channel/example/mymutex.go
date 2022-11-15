package main

import (
	"fmt"
	"time"
)

type myMutex chan struct{}

func NewMyMutex() myMutex {
	ch := make(chan struct{}, 1)
	return ch
}

func (m *myMutex) Lock() {
	(*m) <- struct{}{}
}

func (m *myMutex) UnLock() {
	t := time.NewTimer(1 * time.Second)
	select {
	case <-(*m):
		fmt.Println("释放锁成功。。。")
	case <-t.C:
		fmt.Println("释放锁超时，应该小于等于5秒，。。。")
	case <-time.After(2 * time.Second):
		fmt.Println("超时。。。。。。")
	}
}

func main() {
	num := 1
	m := NewMyMutex()
	m.Lock()
	i := sum(num)
	fmt.Println("sum:", i)
	defer m.UnLock()
}

func sum(a int) int {
	a = a * 2
	return a
}
