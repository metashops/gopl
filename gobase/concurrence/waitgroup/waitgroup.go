package main

import (
	`fmt`
	`sync`
)

type Person struct {
	mu     sync.RWMutex
	salary int
	level  int
}

func (p *Person) promote(wg *sync.WaitGroup) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.salary += 1000
	p.level += 1
	fmt.Printf("==salary: %d, level: %d\n", p.salary, p.level)
	wg.Done()
}

func (p *Person) printPerson(wg *sync.WaitGroup) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fmt.Println("salary:", p.salary)
	fmt.Println("level:", p.level)
	wg.Done()
}

func main() {
	p := Person{salary: 1000, level: 1}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go p.promote(&wg)
	go p.printPerson(&wg)
	go p.printPerson(&wg)
	wg.Wait()
}
