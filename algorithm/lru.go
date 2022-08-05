package algorithm

import "container/list"

// 双向列表 + 哈希表

type LRUCache struct {
	Cap   int
	Len   int
	cache map[int]*list.Element
	list  *list.List
}

type pair struct {
	key int
	val int
}

func NewLRUCache(cap int) LRUCache {
	return LRUCache{
		Cap:   cap,
		list:  list.New(),
		cache: make(map[int]*list.Element),
	}
}

func (lru *LRUCache) Get(key int) int {
	if e, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(e)
		return e.Value.(pair).val
	}
	return -1
}

func (lru *LRUCache) Put(key, val int) int {
	if e, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(e)
		oldVal := e.Value.(pair).val
		e.Value = pair{key: key, val: val}
		return oldVal
	} else {
		if lru.list.Len() >= lru.Cap {
			delete(lru.cache, lru.list.Back().Value.(pair).key)
			lru.list.Remove(lru.list.Back())
		}
		lru.list.PushFront(pair{key: key, val: val})
		lru.cache[key] = lru.list.Front()
		return pair{}.val
	}
}
