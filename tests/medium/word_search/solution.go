package word_search

func exist(board [][]byte, word string) bool {
	rows, columns := len(board), len(board[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(row, column int, word string, seen map[int]bool) bool
	dfs = func(row, column int, word string, seen map[int]bool) bool {
		if row < 0 || row >= rows || column < 0 || column >= columns {
			return false
		}

		key := row*columns + column
		if seen[key] {
			return false
		}

		if word[0] != board[row][column] {
			return false
		}

		if len(word) == 1 {
			return true
		}

		seen[key] = true
		nextWord := word[1:]

		for _, dir := range directions {
			newRow, newCol := row+dir[0], column+dir[1]
			if dfs(newRow, newCol, nextWord, seen) {
				return true
			}
		}

		delete(seen, key)
		return false
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			if board[row][col] != word[0] {
				continue
			}
			if dfs(row, col, word, make(map[int]bool)) {
				return true
			}
		}
	}
	return false
}
