package kth_smallest_element_in_a_bst

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -1

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/
func kthSmallest(root *TreeNode, k int) int {
	found := 0
	if root == nil || k == 0 {
		return found
	}
	n := 0
	var inorderTraversal func(node *TreeNode)
	inorderTraversal = func(node *TreeNode) {
		if n >= k {
			return
		}
		if node.Left != nil {
			inorderTraversal(node.Left)
		}
		n += 1
		if n == k {
			found = node.Val
			return
		}
		if node.Right != nil {
			inorderTraversal(node.Right)
		}
	}
	inorderTraversal(root)
	return found
}
