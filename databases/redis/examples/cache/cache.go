package cache

type Cache interface {
	Get(k interface{})
	Put(k, v interface{})
}
