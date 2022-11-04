package main

import (
	`fmt`
	`sync`
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
	fmt.Printf("salary: %d, level: %d\n", p.salary, p.level)
}

func (p *Person) printPerson() {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println(p.salary)
	fmt.Println(p.level)
}

func main() {
	p := Person{salary: 1000, level: 1}
	for i := 0; i < 10; i++ {
		go p.printPerson()

	}
	for i := 0; i < 10; i++ {
		go p.promote()

	}
	time.Sleep(time.Second)
}
