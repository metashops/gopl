需求：
* 整个程序运行过程中，代码只执行一次
* 用来进行一些初始化的操作

思路1：
* 通过CAS改值，找一个变量记录，从 0 变成 1 就不再执行。
* 优点算法非常简单
* 缺点：多个协程竞争CAS改值会造成大量协程阻塞，导致性能问题

```go
package main

import (
	`fmt`
	`sync`
	`sync/atomic`
	`time`
)

type Person struct {
	mu     sync.RWMutex
	salary int
	level  int
}

func (p *Person) promote() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.salary += 1000
	p.level += 1
	fmt.Printf("==salary: %d, level: %d\n", p.salary, p.level)
}

func main() {
	p := Person{salary: 1000, level: 1}

	once := sync.Once{}

	o := int32(0)
	atomic.CompareAndSwapInt32(&o, 0, 1) // 从 0 变成 1
	time.Sleep(time.Second)
}

```

思路2：
* 竞强一个 mutex，抢不到的陷入 sema 休眠
* 抢到的执行代码，改值，释放锁
* 其他协程唤醒后判断值已经修改，就不需要再去操作了

once 总结：
* 先判断是否已经改值了
* 没改，尝试获取锁
* 获取到锁的协程执行业务，改值，解锁
* 冲突协程唤醒后直接返回

场景：
* 只执行一次的代码或一段逻辑，如初始化
