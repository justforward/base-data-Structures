// 实现一个带有ttl的LRU

package main

import (
	"container/list"
	"time"
)

type LRUCacheWithTTL struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	ttl      time.Duration
}

type entrys struct {
	key   string
	value interface{}
	ttl   time.Time
}

func NewLRUCacheWithTTL(capacity int, ttl time.Duration) *LRUCacheWithTTL {
	return &LRUCacheWithTTL{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		ttl:      ttl,
	}
}

func (c *LRUCacheWithTTL) Get(key string) (interface{}, bool) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		if time.Now().Before(elem.Value.(*entrys).ttl) {
			return elem.Value.(*entry).value, true
		} else {
			c.Remove(key)
		}
	}
	return nil, false
}

func (c *LRUCacheWithTTL) Put(key string, value interface{}) {
	if elem, ok := c.cache[key]; ok {
		c.list.MoveToFront(elem)
		elem.Value.(*entrys).value = value
		elem.Value.(*entrys).ttl = time.Now().Add(c.ttl)
	} else {
		if c.capacity == 0 {
			return
		} else if len(c.cache) >= c.capacity {
			oldestKey := c.list.Back().Value.(*entrys).key
			c.Remove(oldestKey)
		}
		entry := &entrys{
			key:   key,
			value: value,
			ttl:   time.Now().Add(c.ttl),
		}
		elem := c.list.PushFront(entry)
		c.cache[key] = elem
	}
}

func (c *LRUCacheWithTTL) Remove(key string) {
	elem, ok := c.cache[key]
	if !ok {
		return
	}
	c.list.Remove(elem)
	delete(c.cache, key)
}
