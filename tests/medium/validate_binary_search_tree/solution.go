package validate_binary_search_tree

//lint:file-ignore ST1001 _

import (
	"math"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = math.MinInt32 - 1

/*
*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left
subtree

	of a node contains only nodes with keys less than the node's key.

The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var preOrderTraversal func(node *TreeNode)
	lastSeen := null
	valid := true
	preOrderTraversal = func(node *TreeNode) {
		if !valid {
			return
		}
		if node.Left != nil {
			preOrderTraversal(node.Left)
		}
		if node.Val <= lastSeen { // 不能是等於
			valid = false
		}
		lastSeen = node.Val
		if node.Right != nil {
			preOrderTraversal(node.Right)
		}
	}
	preOrderTraversal(root)
	return valid
}
