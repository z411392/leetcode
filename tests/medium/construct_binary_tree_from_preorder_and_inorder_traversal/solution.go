package construct_binary_tree_from_preorder_and_inorder_traversal

//lint:file-ignore ST1001 _

import (
	"slices"

	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

/*
*
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solutions/4289972/golang-short-and-easy-to-understand-solution-with-explaination/
	if len(preorder) == 0 {
		return nil
	}
	// preorder[0] 可以告訴我們 root 是哪個節點
	// inorder 可以讓我們分出 root 的左半邊和右半邊（知道左右半邊的各自的節點數）
	index := slices.Index(inorder, preorder[0])
	inorderLeftPart := inorder[:index]
	inorderRightPart := inorder[index+1:]
	preorderLeftPart := preorder[1 : index+1] // 同樣數目但略過第一個
	preorderRightPart := preorder[index+1:]
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorderLeftPart, inorderLeftPart),
		Right: buildTree(preorderRightPart, inorderRightPart),
	}
}
