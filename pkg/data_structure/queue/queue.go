package queue

import (
	"sync"
)

// Node 代表佇列中的一個節點
type Node struct {
	value interface{}
	next  *Node
}

// Queue 是一個執行緒安全的佇列實作
type Queue struct {
	head *Node
	tail *Node
	size int
	rwmu sync.RWMutex
}

// NewQueue 建立一個新的 Queue
func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Enqueue 將元素加入佇列尾端
func (queue *Queue) Enqueue(value interface{}) {
	queue.rwmu.Lock()
	defer queue.rwmu.Unlock()

	newNode := &Node{
		value: value,
		next:  nil,
	}

	if queue.tail == nil {
		queue.head = newNode
		queue.tail = newNode
	} else {
		queue.tail.next = newNode
		queue.tail = newNode
	}
	queue.size++
}

// Dequeue 從佇列前端移除並返回元素
func (queue *Queue) Dequeue() (interface{}, bool) {
	queue.rwmu.Lock()
	defer queue.rwmu.Unlock()

	if queue.head == nil {
		return nil, false
	}

	value := queue.head.value
	queue.head = queue.head.next
	queue.size--

	if queue.head == nil {
		queue.tail = nil
	}

	return value, true
}

// Peek 查看佇列前端的元素但不移除
func (queue *Queue) Peek() (interface{}, bool) {
	queue.rwmu.RLock()
	defer queue.rwmu.RUnlock()

	if queue.head == nil {
		return nil, false
	}

	return queue.head.value, true
}

// Size 返回佇列中的元素數量
func (queue *Queue) Size() int {
	queue.rwmu.RLock()
	defer queue.rwmu.RUnlock()

	return queue.size
}

// IsEmpty 檢查佇列是否為空
func (queue *Queue) IsEmpty() bool {
	return queue.Size() == 0
}

// ToStream 返回佇列的一個複製，不影響原佇列
func (queue *Queue) ToStream() []interface{} {
	queue.rwmu.RLock()
	defer queue.rwmu.RUnlock()

	stream := make([]interface{}, queue.size)
	current := queue.head
	i := 0

	for current != nil {
		stream[i] = current.value
		current = current.next
		i++
	}

	return stream
}
