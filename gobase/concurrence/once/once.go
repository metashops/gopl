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

func (p *Person) printPerson() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println("salary:", p.salary)
	fmt.Println("level:", p.level)
}

func main() {
	p := Person{salary: 1000, level: 1}

	once := sync.Once{}

	o := int32(0)
	atomic.CompareAndSwapInt32(&o, 0, 1) // 从 0 变成 1
	go once.Do(p.promote)
	go once.Do(p.promote)
	go once.Do(p.promote)
	time.Sleep(time.Second)
}
