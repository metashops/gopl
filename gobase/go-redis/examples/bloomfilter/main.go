package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/bits-and-blooms/bitset"
)

type BloomFilter struct {
	bitmap *bitset.BitSet // 位图
	k      uint           // 哈希函数个数
}

func NewBloomFilter(n uint, k uint) *BloomFilter {
	bitmap := bitset.New(n * 10)
	return &BloomFilter{
		bitmap: bitmap,
		k:      k,
	}
}

func (b *BloomFilter) Add(element string) {
	for i := uint(0); i < b.k; i++ {
		hash := hash(i, element)
		b.bitmap.Set(hash)
	}
}

func (b *BloomFilter) Check(element string) bool {
	for i := uint(0); i < b.k; i++ {
		hash := hash(i, element)
		if !b.bitmap.Test(hash) {
			return false
		}
	}
	return true
}

// 哈希函数，使用SHA-256哈希算法
func hash(n uint, element string) uint {
	if n == 0 {
		n = 1000
	}
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%s", n, element)))
	return uint(hash[0]) % (n * 10 * 8)
}

func main() {
	b := NewBloomFilter(1000, 5)
	b.Add("hello")
	b.Add("world")
	b.Add("world1")
	b.Add("world2")
	fmt.Println(b.Check("hello")) // true
	fmt.Println(b.Check("world")) // true
	fmt.Println(b.Check("foo"))   // true   // false
}
