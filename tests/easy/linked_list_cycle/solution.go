package best_time_to_buy_and_sell_stock

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/
func hasCycle(head *ListNode) bool {
	node := head
	if node == nil {
		return false
	}
	var seen = map[*ListNode]bool{}
	for {
		if _, ok := seen[node]; ok {
			return true
		}
		seen[node] = true
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return false
}
