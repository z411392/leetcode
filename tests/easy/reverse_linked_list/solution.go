package reverse_linked_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
	. "github.com/z411392/leetcode/pkg/data_structure/stack"
)

/*
Given the head of a singly linked list, reverse the list, and return the reversed list.
*/
func reverseList(head *ListNode) *ListNode {
	stack := NewStack[int]()
	streamCh := ConvertLinkedListToStream(head)
	for value := range streamCh {
		stack.Push(value)
	}
	var list *ListNode = nil
	var lastNode *ListNode = nil
	for {
		if stack.Size() == 0 {
			break
		}
		value, _ := stack.Pop()
		node := &ListNode{
			Val:  value,
			Next: nil,
		}
		if list == nil {
			list = node
		} else {
			lastNode.Next = node
		}
		lastNode = node
	}
	return list
}
