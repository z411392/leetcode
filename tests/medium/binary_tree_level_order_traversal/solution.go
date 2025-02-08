package binary_tree_level_order_traversal

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -1001

/*
*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/
func levelOrder(root *TreeNode) [][]int {
	levelsData := [][]int{}
	if root == nil {
		return levelsData
	}
	queue := [](*TreeNode){root}
	var nodes [](*TreeNode)
	for {
		nodes = queue
		if len(nodes) == 0 {
			break
		}
		queue = [](*TreeNode){}
		levelData := []int{}
		for _, node := range nodes {
			levelData = append(levelData, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelsData = append(levelsData, levelData)
		nodes = queue
	}
	return levelsData
}
