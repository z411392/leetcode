package maximum_depth_of_binary_tree

import (
	//lint:ignore ST1001 _
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -101

func depth(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	level += 1
	left := level
	right := level
	if root.Left != nil {
		left = depth(root.Left, level)
	}
	if root.Right != nil {
		right = depth(root.Right, level)
	}
	if left >= right {
		return left
	}
	return right
}

/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/
func maxDepth(root *TreeNode) int {
	return depth(root, 0)
}
