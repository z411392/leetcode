package swap_nodes_in_pairs

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func swapPairs(head *ListNode) *ListNode {
	// https://leetcode.com/problems/swap-nodes-in-pairs/solutions/300006/go-0-ms-faster-than-100-00-easy-code-recursion/
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}
