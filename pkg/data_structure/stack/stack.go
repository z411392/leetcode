package stack

import (
	"sync"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	top  *Node[T]
	mu   sync.RWMutex
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newNode := &Node[T]{value: item, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var zeroValue T
	if s.top == nil {
		return zeroValue, false
	}

	item := s.top.value
	s.top = s.top.next
	s.size--
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var zeroValue T
	if s.top == nil {
		return zeroValue, false
	}

	return s.top.value, true
}

func (s *Stack[T]) ToStream() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]T, s.size)
	current := s.top
	for i := 0; i < s.size; i++ {
		result[i] = current.value
		current = current.next
	}
	return result
}

func (s *Stack[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.size
}
