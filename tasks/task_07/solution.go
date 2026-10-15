package main

import (
	"container/list"
	"sync"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRUCache[K comparable, V any] struct {
	mu sync.Mutex

	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		items:    make(map[K]*list.Element),
	}
}

func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	if c.capacity <= 0 {
		return value, false
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.items[key]; found {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}

	return value, false
}

func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.items[key]; found {
		elem.Value.(*entry[K, V]).value = value
		return
	}

	if c.ll.Len() >= c.capacity {
		backElem := c.ll.Back()
		if backElem != nil {
			c.ll.Remove(backElem)
			delete(c.items, backElem.Value.(*entry[K, V]).key)
		}
	}

	newEntry := &entry[K, V]{key: key, value: value}
	elem := c.ll.PushFront(newEntry)
	c.items[key] = elem
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
