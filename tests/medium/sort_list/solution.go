package sort_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func sortList(head *ListNode) *ListNode {
	// Base cases
	if head == nil || head.Next == nil {
		return head
	}

	// Find middle point using slow and fast pointers
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Split the list into two halves
	prev.Next = nil

	// Recursively sort both halves
	left := sortList(head)
	right := sortList(slow)

	// Merge the sorted halves
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	// Append remaining nodes
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
