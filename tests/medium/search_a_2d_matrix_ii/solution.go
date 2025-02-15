package search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	row, col := 0, n-1 // 從右上角開始
	for {
		if col < 0 {
			break
		}
		if row >= m {
			break
		}
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col -= 1 // 向左移動
			continue
		}
		row += 1 // 向下移動
	}

	return false
}
