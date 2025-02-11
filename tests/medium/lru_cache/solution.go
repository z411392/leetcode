package lru_cache

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
int get(int key) Return the value of the key if the key exists, otherwise return -1.
void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
The functions get and put must each run in O(1) average time complexity.
*/
type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	// 預分配 capacity 大小的 map
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		// 直接在這裡內聯 moveToEnd 的邏輯
		// 1. 從當前位置移除
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// 2. 移動到尾部
		lastNode := this.tail.Prev
		lastNode.Next = node
		node.Prev = lastNode
		node.Next = this.tail
		this.tail.Prev = node

		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		// 更新值
		node.Value = value

		// 直接內聯 moveToEnd 的邏輯
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		lastNode := this.tail.Prev
		lastNode.Next = node
		node.Prev = lastNode
		node.Next = this.tail
		this.tail.Prev = node
		return
	}

	// 創建新節點
	newNode := &Node{
		Key:   key,
		Value: value,
	}
	this.cache[key] = newNode

	// 直接加到尾部
	lastNode := this.tail.Prev
	lastNode.Next = newNode
	newNode.Prev = lastNode
	newNode.Next = this.tail
	this.tail.Prev = newNode

	// 檢查容量
	if len(this.cache) > this.capacity {
		// 移除最久未使用的節點
		lru := this.head.Next
		this.head.Next = lru.Next
		lru.Next.Prev = this.head
		delete(this.cache, lru.Key)
	}
}
