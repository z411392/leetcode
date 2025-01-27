package palindrome_linked_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
	. "github.com/z411392/leetcode/pkg/data_structure/stack"
)

/*
Given the head of a singly linked list, return true if it is a palindrome or false otherwise.
*/
func isPalindrome(head *ListNode) bool {
	stack := NewStack[int]()
	streamCh := ConvertLinkedListToStream(head)
	for value := range streamCh {
		stack.Push(value)
	}
	stackMidSize := stack.Size() / 2
	for node := head; node != nil; node = node.Next {
		if stack.Size() < stackMidSize {
			return true
		}
		value, _ := stack.Pop()
		if node.Val != value {
			return false
		}
	}
	return true
}
