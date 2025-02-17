package odd_even_linked_list

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func oddEvenList(head *ListNode) *ListNode {
	node := head
	if node == nil {
		return nil
	}
	oddHead := &ListNode{Val: node.Val}
	prevOddNode := oddHead

	node = node.Next
	if node == nil {
		return oddHead
	}
	evenHead := &ListNode{Val: node.Val}
	prevEvenNode := evenHead
	isEven := true
	for {
		node = node.Next
		if node == nil {
			break
		}
		isEven = !isEven
		// fmt.Printf("%v %v\n", node.Val, isEven)
		newNode := &ListNode{Val: node.Val}
		if isEven {
			prevEvenNode.Next = newNode
			prevEvenNode = newNode
		} else {
			prevOddNode.Next = newNode
			prevOddNode = newNode
		}
	}
	prevOddNode.Next = evenHead
	return oddHead
}
