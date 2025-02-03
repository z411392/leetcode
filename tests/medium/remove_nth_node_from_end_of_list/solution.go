package remove_nth_node_from_end_of_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var node *ListNode = head
	seen := [](*ListNode){}
	for {
		if node == nil {
			break
		}
		seen = append(seen, node)
		node = node.Next
	}
	index := len(seen) - n
	// 先排除 index 不存在的情況
	if index < 0 || index >= len(seen) {
		return head
	}
	// 如果只有單一節點直接回傳 nil
	if len(seen) == 1 {
		return nil
	}
	// index 位在開頭的情形（直接把第二個節點當作 head 回傳）
	if index == 0 {
		return seen[1]
	}
	// index 位在結尾的情形（將倒數第二個節點也就是 index - 2 的 Next 設為 nil）
	if index == len(seen)-1 {
		seen[index-1].Next = nil
		return head
	}
	// 否則將 index - 1 的 Next 設為 index + 1
	previous := index - 1
	next := index + 1
	seen[previous].Next = seen[next]
	return head
}
