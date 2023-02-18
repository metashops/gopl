package cache

import (
	"container/list"
	"sync"
	"sync/atomic"
)

/*
	思路：可以使用list加map实现LRU cache
	具体思路为:写入时。先从map中查询,如果能查询,如果能查询到值,则将该值的在List中移动到最前面；
			 如果查询不到值,则判断当前map是否到达最大值,如果到达最大值则移除List最后面的值,同时删除map中的值；
			 如果map容量未达最大值,则写入map,同时将值放在List最前面。
			 读取时。从map中查询,如果能查询到值,则直接将List中该值移动到最前面,返回查询结果.为保证并发安全,需要引入读写锁.
			 另外,存在读取List中内容反差map的情况,因为声明一个容器对象同时保存key以及value,
			 List中以及map中存储的都是容器对象的引用.引入原子对象对命中数以及未命中数等指标进行统计
			 原文：https://blog.csdn.net/text2203/article/details/128640113
*/

// An AtomicInt is an int64 to be accessed atomically.
type AtomicInt int64

// MemCache is an LRU cache. It is safe for concurrent access.
type MemCache struct {
	m           sync.RWMutex
	maxItemSize int
	cacheList   *list.List
	cache       map[interface{}]*list.Element
	hits, gets  AtomicInt
}

type entry struct {
	key   interface{}
	value interface{}
}

// NewMemCache If maxItemSize is zero, the cache has no limit.
// if maxItemSize is not zero, when cache's size beyond maxItemSize,start to swap
func NewMemCache(maxItemSize int) *MemCache {
	return &MemCache{
		maxItemSize: maxItemSize,
		cacheList:   list.New(),
		cache:       make(map[interface{}]*list.Element),
	}
}

// Get value with key
func (c *MemCache) Get(key string) (interface{}, bool) {
	c.m.RLock()
	defer c.m.RUnlock()
	c.gets.Add(1) // 如果读取到值,移动在List中位置,并返回value
	if elem, hit := c.cache[key]; hit {
		c.hits.Add(1)
		c.cacheList.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// Set a value with key
func (c *MemCache) Set(key string, value interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.cacheList = list.New()
	}

	// 判断是否在map中,如果在map中,则将value从list中移动到前面.
	// 如果键不存在，ok 的值为 false，elem 的值为该类型的零值
	if elem, ok := c.cache[key]; ok {
		c.cacheList.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	// //如果不再map中,将值存到List最前面
	elem := c.cacheList.PushFront(&entry{key: key, value: value})
	c.cache[key] = elem // 判断是否到达容量限制,到达容量限制时删除List中最后面的值.
	if c.maxItemSize != 0 && c.cacheList.Len() > c.maxItemSize {
		c.RemoveOldest()
	}
}

// RemoveOldest remove the oldest key
func (c *MemCache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.cacheList.Back()
	if ele != nil {
		c.cacheList.Remove(ele)
		key := ele.Value.(*entry).key
		delete(c.cache, key)
	}
}

// Add atomically adds n to i.
func (i *AtomicInt) Add(n int64) {
	atomic.AddInt64((*int64)(i), n)
}

// Get atomically gets the value of i.
func (i *AtomicInt) Get() int64 {
	return atomic.LoadInt64((*int64)(i))
}
