package lru_cache

import "slices"

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/
type LRUCache struct {
	Capacity int
	Values   map[int]int
	Keys     []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Values:   map[int]int{},
		Keys:     []int{},
	}
}

func (this *LRUCache) touch(key int) {
	index := slices.Index(this.Keys, key)
	if index > -1 {
		this.Keys = append(this.Keys[:index], this.Keys[index+1:]...)
	}
	this.Keys = append(this.Keys, key)
}

func (this *LRUCache) Get(key int) int {
	value, exists := this.Values[key]
	if !exists {
		return -1
	}
	this.touch(key)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	this.Values[key] = value
	this.touch(key)
	if len(this.Keys) > this.Capacity {
		key := this.Keys[0]
		delete(this.Values, key)
		this.Keys = this.Keys[1:]

	}
}
