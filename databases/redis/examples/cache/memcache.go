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

// 节点插入头部，在Head节点后
func (l *LRUCache) attach(n *Node) {
	n.Prev = l.Head
	n.Next = l.Head.Next

	l.Head.Next = n
	n.Next.Prev = n
}

func (l *LRUCache) Put(key, value interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	// 不包含
	if v, ok := l.Map[key]; ok {
		v.V = value
		l.detach(v)
		l.attach(v)

		return
	}

	var n *Node
	if len(l.Map) >= l.Capacity {
		// 说明已经达到最大容量
		n = l.Tail.Prev
		l.detach(n)
		delete(l.Map, n.K)

	} else {
		n = &Node{}
	}
	n.K = key
	n.V = value
	l.Map[n.K] = n
	l.attach(n)
}

func (l *LRUCache) Get(key interface{}) interface{} {
	l.m.RLock()
	defer l.m.RUnlock()

	if v, ok := l.Map[key]; ok {
		// 将节点放到最前面
		l.detach(v)
		l.attach(v)
		return v.V
	}
	return interface{}(-1)
}
