package add_two_numbers

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var lastNode *ListNode = nil
	stream1 := ConvertLinkedListToStream(l1)
	stream2 := ConvertLinkedListToStream(l2)
	carry := false
	for {
		value1, ok1 := <-stream1
		value2, ok2 := <-stream2
		value := value1 + value2
		if carry {
			value += 1
			carry = false
		}
		// 不能靠 ok1, ok2 都是 false 就跳掉（還有 carry 是 true 的情形）
		// 也不能因為 value 是 0 就直接跳掉，只有一位時還要顯示這個零
		if !ok1 && !ok2 && value == 0 {
			if head == nil {
				head = &ListNode{Val: value}
			}
			break
		}
		if value >= 10 {
			carry = true
			value -= 10
		}
		node := &ListNode{Val: value}
		if head == nil {
			head = node
		} else {
			lastNode.Next = node
		}
		lastNode = node
	}
	return head
}
