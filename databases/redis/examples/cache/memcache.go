package cache

import (
	"sync"
)

// Node 双向链表
type Node struct {
	K    interface{}
	V    interface{}
	Prev *Node
	Next *Node
}

// LRUCache 数据结构:双向链表+哈希表
type LRUCache struct {
	Capacity   int
	Head, Tail *Node
	Map        map[interface{}]*Node
	m          sync.RWMutex
}

func (l *LRUCache) NewLRUCache(capacity int) *LRUCache {
	l.Capacity = capacity
	l.Head = &Node{}
	l.Tail = &Node{}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	l.Head.Prev = nil
	l.Tail.Next = nil

	l.Map = make(map[interface{}]*Node)

	return l
}

// 分离节点
func (l *LRUCache) detach(n *Node) {
	n.Prev.Next = n.Next // 该节点的前一个节点的下一个节点，指向next的节点
	n.Next.Prev = n.Prev
}

func (l *LRUCache) Put(key, value interface{}) {}
func (l *LRUCache) Get(key interface{}) {
	//
}
