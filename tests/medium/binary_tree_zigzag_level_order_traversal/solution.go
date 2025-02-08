package binary_tree_zigzag_level_order_traversal

//lint:file-ignore ST1001 _

import (
	. "github.com/z411392/leetcode/pkg/data_structure/tree"
)

const null = -101

/*
*
Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
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
		dir := len(levelsData) % 2
		for _, node := range nodes {
			if dir == 0 {
				levelData = append(levelData, node.Val)
			} else {
				levelData = append([]int{node.Val}, levelData...)
			}
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
