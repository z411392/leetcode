package set_matrix_zeroes

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
*/
func setZeroes(matrix [][]int) {
	rows := map[int]bool{}
	columns := map[int]bool{}
	for row, cells := range matrix {
		for column, cell := range cells {
			if cell != 0 {
				continue
			}
			rows[row] = true
			columns[column] = true
		}
	}
	// 假設 rows=4 columns=3 組在一起跑是 4x3=12，但分開跑是 4+3=7
	for row := range rows {
		for column := range matrix[row] {
			matrix[row][column] = 0
		}
	}
	for column := range columns {
		for row := range matrix {
			matrix[row][column] = 0
		}
	}
}
