package diameter_of_binary_tree

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -101

// https://leetcode.com/problems/diameter-of-binary-tree/solutions/4523542/go-solution-great-explanation-and-full-description/
var maxD int

/*
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/
func diameterOfBinaryTree(root *TreeNode) int {
	maxD = 0
	find(root)
	return maxD
}

func find(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := find(root.Left)
	right := find(root.Right)
	localMax := left + right
	maxD = max(maxD, localMax)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
