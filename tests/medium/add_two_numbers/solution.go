package add_two_numbers

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

// https://leetcode.com/problems/add-two-numbers/solutions/6257320/100-beats-go

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var lastNode *ListNode = nil
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		if head == nil {
			head = node
		} else {
			lastNode.Next = node
		}
		lastNode = node
	}
	return head
}
