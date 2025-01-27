package invert_binary_tree

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -101

/*
Given the root of a binary tree, invert the tree, and return its root.
*/
func invertTree(root *TreeNode) *TreeNode {
	// https://leetcode.com/problems/invert-binary-tree/solutions/1695693/go-4-lines-solution-clean-solution-in-go-golang-recursive-0ms-100/
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
