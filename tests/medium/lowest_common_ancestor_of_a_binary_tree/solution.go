package lowest_common_ancestor_of_a_binary_tree

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

// const null = math.MinInt32

/*
Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var ancestorsOf func(node *TreeNode, target *TreeNode) []*TreeNode
	ancestorsOf = func(node *TreeNode, target *TreeNode) []*TreeNode {
		if node == target {
			return []*TreeNode{node}
		}
		if node.Left != nil {
			ancestors := ancestorsOf(node.Left, target)
			if ancestors != nil {
				return append(ancestors, node)
			}
		}
		if node.Right != nil {
			ancestors := ancestorsOf(node.Right, target)
			if ancestors != nil {
				return append(ancestors, node)
			}
		}
		return nil
	}
	ptrs1 := ancestorsOf(root, p)
	// fmt.Printf("%v\n", ptrs1)
	ptrs2 := ancestorsOf(root, q)
	// fmt.Printf("%v\n", ptrs2)
	var lca *TreeNode = nil
	for {
		if len(ptrs1) == 0 || len(ptrs2) == 0 {
			break
		}
		ptr1 := ptrs1[len(ptrs1)-1]
		ptr2 := ptrs2[len(ptrs2)-1]
		if ptr1 == ptr2 {
			ptrs1 = ptrs1[:len(ptrs1)-1]
			ptrs2 = ptrs2[:len(ptrs2)-1]
			lca = ptr1
			continue
		}
		break
	}
	return lca
}
