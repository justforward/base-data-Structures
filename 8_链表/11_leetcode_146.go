package main

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	cap       int
	list      *list.List
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, list: list.New(), keyToNode: map[int]*list.Element{}}
}

func (this *LRUCache) Get(key int) int {
	node := this.keyToNode[key]
	if node == nil {
		return -1
	}
	this.list.MoveToFront(node)
	return node.Value.(entry).value
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.keyToNode[key]; node != nil {
		node.Value = entry{key, value}
		this.list.MoveToFront(node)
		return
	}
	this.keyToNode[key] = this.list.PushFront(entry{key, value})
	if len(this.keyToNode) > this.cap {
		delete(this.keyToNode, this.list.Remove(this.list.Back()).(entry).key)
	}
}
