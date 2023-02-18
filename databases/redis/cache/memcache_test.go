package cache

import (
	"testing"
)

type simpleStruct struct {
	int
	string
}

type complexStruct struct {
	int
	simpleStruct
}

var getTests = []struct {
	name       string
	keyToAdd   string
	keyToGet   string
	expectedOk bool
}{
	{"string_hit", "myKey", "myKey", true},
	{"string_miss", "myKey", "nonsense", false},
}

func TestSet(t *testing.T) {
	var cache Cache
	cache = NewMemCache(0)
	values := []string{"test1", "test2", "test3"}
	key := "key1"
	for _, v := range values {
		cache.Set(key, v)
		val, ok := cache.Get(key)
		if !ok {
			t.Fatalf("expect key:%v ,value:%v", key, v)
		} else if ok && val != v {
			t.Fatalf("expect value:%v, get value:%v", key, v, val)
		}
		t.Logf("value:%v ", val)
	}
}

func TestGet(t *testing.T) {
	var cache Cache
	cache = NewMemCache(0)
	for _, tt := range getTests {
		cache.Set(tt.keyToAdd, 1234)
		val, ok := cache.Get(tt.keyToGet)

		if ok != tt.expectedOk {
			t.Fatalf("%s: val:%v cache hit = %v; want %v", tt.name, val, ok, !ok)
		} else if ok && val != 1234 {
			t.Fatalf("%s expected get to return 1234 but got %v", tt.name, val)
		}

	}
}
