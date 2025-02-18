package kth_smallest_element_in_a_sorted_matrix

func kthSmallest(matrix [][]int, k int) int {
	min, max := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]
	// 找小於某個中間值的元素個數
	// 如果少於 k

	for min != max {
		mid := min + (max-min)/2
		if countSmallerNumbers(matrix, mid) < k {
			min = mid + 1
		} else {
			max = mid
		}
	}

	return min
}

func countSmallerNumbers(matrix [][]int, mid int) int {
	// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/solutions/2369817/golang-solution-binary-search/
	count, row, col := 0, 0, len(matrix[0])-1
	for {
		if row >= len(matrix) {
			break
		}
		if col < 0 {
			break
		}
		value := matrix[row][col]
		if value <= mid {
			count += col + 1 // 一次加完所有小於等於 mid 的元素個數
			row++            // 然後往下移動
		} else {
			col-- // 否則往左移動到底
		}
		// 雖然走的是 matrix 但因為一次只走某個方向所以步數是少的
	}

	return count
}
