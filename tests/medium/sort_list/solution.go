package sort_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

type MinHeap struct {
	nodes []*ListNode
}

func NewMinHeap() *MinHeap {
	return &MinHeap{nodes: make([]*ListNode, 0)}
}

func (h *MinHeap) Len() int {
	return len(h.nodes)
}

func (h *MinHeap) swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
}

func (h *MinHeap) less(i, j int) bool {
	return h.nodes[i].Val < h.nodes[j].Val
}

func (h *MinHeap) Push(node *ListNode) {
	// 1. 加入到尾部
	h.nodes = append(h.nodes, node)

	// 2. 上浮操作
	h.upHeap(len(h.nodes) - 1)
}

func (h *MinHeap) Pop() *ListNode {
	if len(h.nodes) == 0 {
		return nil
	}

	// 1. 保存要返回的根節點
	root := h.nodes[0]

	// 2. 將最後一個元素移到根位置
	lastIdx := len(h.nodes) - 1
	h.nodes[0] = h.nodes[lastIdx]
	h.nodes = h.nodes[:lastIdx]

	// 3. 下沉操作
	if len(h.nodes) > 0 {
		h.downHeap(0)
	}

	return root
}

func (h *MinHeap) upHeap(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || !h.less(i, parent) {
			break
		}
		h.swap(i, parent)
		i = parent
	}
}

func (h *MinHeap) downHeap(i int) {
	for {
		smallest := i
		left := 2*i + 1
		right := 2*i + 2

		if left < len(h.nodes) && h.less(left, smallest) {
			smallest = left
		}
		if right < len(h.nodes) && h.less(right, smallest) {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.swap(i, smallest)
		i = smallest
	}
}

/*
Given the head of a linked list, return the list after sorting it in ascending order.
*/
func sortList(head *ListNode) *ListNode {
	heap := NewMinHeap()
	node := head
	for {
		if node == nil {
			break
		}
		heap.Push(node)
		node = node.Next
	}
	var newHead *ListNode = nil
	var lastNode *ListNode = nil
	for {
		if heap.Len() == 0 {
			break
		}
		node := heap.Pop()
		node.Next = nil
		if newHead == nil {
			newHead = node
		} else {
			lastNode.Next = node
		}
		lastNode = node
	}
	return newHead
}
