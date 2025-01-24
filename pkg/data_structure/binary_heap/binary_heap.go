package binary_heap

import (
	"sync"
)

// Node represents a single node in the binary heap.
type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// CompareFunc is used to determine the order of the heap.
// It should return true if the first argument has higher priority than the second.
type CompareFunc[T any] func(a, b T) bool

// BinaryHeap represents a binary heap.
type BinaryHeap[T any] struct {
	root    *Node[T]
	compare CompareFunc[T]
	size    int
	rwMux   sync.RWMutex // RWMutex for protecting shared state
}

// NewBinaryHeap creates a new binary heap with the given comparison function.
func NewBinaryHeap[T any](compare CompareFunc[T]) *BinaryHeap[T] {
	return &BinaryHeap[T]{
		compare: compare,
	}
}

// Insert adds a new element to the heap.
func (heap *BinaryHeap[T]) Insert(value T) {
	heap.rwMux.Lock() // Acquire write lock
	defer heap.rwMux.Unlock()

	heap.size++
	newNode := &Node[T]{Value: value}
	if heap.root == nil {
		heap.root = newNode
		return
	}
	heap.root = heap.merge(heap.root, newNode)
}

// Top returns the top element of the heap without removing it.
func (heap *BinaryHeap[T]) Top() (T, bool) {
	heap.rwMux.RLock() // Acquire read lock
	defer heap.rwMux.RUnlock()

	if heap.root == nil {
		var zero T
		return zero, false
	}
	return heap.root.Value, true
}

// Pop removes and returns the top element of the heap.
func (heap *BinaryHeap[T]) Pop() (T, bool) {
	heap.rwMux.Lock()
	defer heap.rwMux.Unlock()

	if heap.root == nil {
		var zero T
		return zero, false
	}

	// Pop the root and reheapify
	value := heap.root.Value
	heap.root = heap.merge(heap.root.Left, heap.root.Right)
	heap.size--
	return value, true
}

// ToStream returns a channel that streams the elements of the heap in order.
func (heap *BinaryHeap[T]) ToStream() <-chan T {
	streamCh := make(chan T)
	tempHeap := heap.clone() // Create a clone to avoid modifying the original heap

	go func() {
		defer close(streamCh)

		// Stream elements from the cloned heap
		for tempHeap.size > 0 {
			value, _ := tempHeap.Pop()
			streamCh <- value
		}
	}()

	return streamCh
}

// clone creates a deep copy of the heap
func (heap *BinaryHeap[T]) clone() *BinaryHeap[T] {
	heap.rwMux.RLock()
	defer heap.rwMux.RUnlock()

	newHeap := NewBinaryHeap(heap.compare)

	// Helper function to recursively clone nodes
	var cloneNode func(*Node[T]) *Node[T]
	cloneNode = func(node *Node[T]) *Node[T] {
		if node == nil {
			return nil
		}
		return &Node[T]{
			Value: node.Value,
			Left:  cloneNode(node.Left),
			Right: cloneNode(node.Right),
		}
	}

	newHeap.root = cloneNode(heap.root)
	newHeap.size = heap.size
	return newHeap
}

// merge combines two subtrees and maintains the heap property.
func (heap *BinaryHeap[T]) merge(a, b *Node[T]) *Node[T] {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	if heap.compare(b.Value, a.Value) {
		a, b = b, a
	}

	a.Right = heap.merge(a.Right, b)
	a.Left, a.Right = a.Right, a.Left // Swap children for balancing

	return a
}

// Size returns the number of elements in the heap.
func (heap *BinaryHeap[T]) Size() int {
	heap.rwMux.RLock() // Acquire read lock
	defer heap.rwMux.RUnlock()

	return heap.size
}
