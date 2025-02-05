package set_matrix_zeroes

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
*/
func setZeroes(matrix [][]int) {
	rows := map[int]bool{}
	columns := map[int]bool{}
	for row, cells := range matrix {
		for column, cell := range cells {
			if cell == 0 {
				rows[row] = true
				columns[column] = true
			}
		}
	}
	for row, cells := range matrix {
		for column := range cells {
			if rows[row] || columns[column] {
				matrix[row][column] = 0
			}
		}
	}
}
