package binary_tree_inorder_traversal

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

func traverseInOrder(root *TreeNode, fun func(value int)) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traverseInOrder(root.Left, fun)
	}
	fun(root.Val)
	if root.Right != nil {
		traverseInOrder(root.Right, fun)
	}
}

func toStream(root *TreeNode, fun func(root *TreeNode, fun func(value int))) <-chan int {
	streamCh := make(chan int)
	go func() {
		defer close(streamCh)
		fun(root, func(value int) {
			streamCh <- value
		})
	}()
	return streamCh
}

const null = -101

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/
func inorderTraversal(root *TreeNode) []int {
	slice := make([]int, 0)
	for value := range toStream(root, traverseInOrder) {
		slice = append(slice, value)
	}
	return slice
}
