package search_a_2d_matrix_ii

import (
	"container/heap"
)

// 定義最小堆的元素
type Cell struct {
	Val         int // 值
	Row, Column int // 在矩陣中的座標
}

// 定義最小堆
type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// DFS + MinHeap 遍歷矩陣
func traverseMatrix(matrix [][]int, forEach func(cell Cell) bool) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// 記錄已訪問的位置
	visited := make(map[[2]int]bool)

	// 方向 (右, 下)
	directions := [][]int{{0, 1}, {1, 0}}

	// 從 (0,0) 開始
	heap.Push(minHeap, Cell{matrix[0][0], 0, 0})
	visited[[2]int{0, 0}] = true

	for minHeap.Len() > 0 {
		cell := heap.Pop(minHeap).(Cell)
		stop := forEach(cell)
		if stop {
			break
		}
		// 向右與向下探索
		for _, dir := range directions {
			newX, newY := cell.Row+dir[0], cell.Column+dir[1]
			if newX < m && newY < n && !visited[[2]int{newX, newY}] {
				heap.Push(minHeap, Cell{matrix[newX][newY], newX, newY})
				visited[[2]int{newX, newY}] = true
			}
		}
	}
}

/*
Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	found := false
	traverseMatrix(matrix, func(cell Cell) bool {
		// fmt.Printf("(%v,%v)=%v\n", cell.Row, cell.Column, cell.Val)
		if cell.Val == target {
			found = true
		}
		return cell.Val >= target
	})
	return found
}
