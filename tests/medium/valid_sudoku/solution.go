package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	// https://bear-1111.medium.com/36-valid-sudoku-leetcode-%E5%BE%88%E9%85%B7%E7%9A%84%E8%BD%89%E6%8F%9B%E7%9F%A9%E9%99%A3-af422c438fcb
	numbersRecordedInRows := new([9][9]bool) // 大小固定的話就不是 slice 而是 array 初始化改用 new 不用 make
	numbersRecordedInColumns := new([9][9]bool)
	numbersRecordedInBoxes := new([9][9]bool)
	for row, cells := range board {
		for column, cell := range cells {
			if cell == '.' {
				continue
			}
			number := int(cell-'0') - 1
			box := row/3*3 + column/3
			if numbersRecordedInRows[row][number] || // 每一列 每一數字 只能被用一次
				numbersRecordedInColumns[column][number] || // 每一行 每一數字 只能被用一次
				numbersRecordedInBoxes[box][number] { // 每一 3x3 矩陣 每一數字 只能被用一次
				return false
			}
			numbersRecordedInRows[row][number] = true
			numbersRecordedInColumns[column][number] = true
			numbersRecordedInBoxes[box][number] = true
		}
	}
	return true
}
