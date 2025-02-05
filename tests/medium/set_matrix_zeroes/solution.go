package set_matrix_zeroes

/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
*/
func setZeroes(matrix [][]int) {
	rows := map[int]bool{}
	columns := map[int]bool{}
	reset := map[[2]int]bool{}
	setRow := func(row int) {
		if rows[row] {
			return
		}
		for column := range matrix[row] {
			reset[[2]int{row, column}] = true
		}
		rows[row] = true
	}
	setColumn := func(column int) {
		if columns[column] {
			return
		}
		for row := 0; row < len(matrix); row += 1 {
			reset[[2]int{row, column}] = true
		}
		columns[column] = true
	}
	for row, cells := range matrix {
		for column, cell := range cells {
			if cell != 0 {
				continue
			}
			setRow(row)
			setColumn(column)
		}
	}
	for position := range reset {
		matrix[position[0]][position[1]] = 0
	}
}
