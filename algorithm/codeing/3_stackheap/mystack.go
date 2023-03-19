package main

import (
	`fmt`
)

func main() {
	stack := Constructor()
	stack.Push(1, 2, 3)
	fmt.Println(stack.queue)
	fmt.Println(stack.Pop())
}

/**
使用队列实现栈
*/

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (s *MyStack) Push(elem ...int) {
	s.queue = append(s.queue, elem...)
}

func (s *MyStack) Pop() int {
	l := len(s.queue) - 1
	for l != 0 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		s.queue = append(s.queue, val)
		l--
	}
	// 弹出
	val := s.queue[0]
	s.queue = s.queue[1:]
	return val
}

func (s *MyStack) Top() int {
	val := s.Pop()
	s.queue = append(s.queue, val)
	return val
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}
