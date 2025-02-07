package word_search

func exist(board [][]byte, word string) bool {
	// 優化：檢查字符頻率
	charCount := make(map[byte]int)
	rows, columns := len(board), len(board[0])

	// 統計 board 中每個字符的出現次數
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			charCount[board[i][j]]++
		}
	}

	// 檢查 word 中的每個字符是否都有足夠的數量
	wordCount := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		wordCount[word[i]]++
		if wordCount[word[i]] > charCount[word[i]] {
			return false
		}
	}

	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(row, column int, index int) bool
	dfs = func(row, column int, index int) bool {
		if index == len(word) {
			return true
		}

		if row < 0 || row >= rows || column < 0 || column >= columns ||
			board[row][column] != word[index] {
			return false
		}

		temp := board[row][column]
		board[row][column] = '#'

		for _, dir := range directions {
			if dfs(row+dir[0], column+dir[1], index+1) {
				return true
			}
		}

		board[row][column] = temp
		return false
	}

	// 優化：從較少出現的字符開始搜索
	firstChar := word[0]
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			if board[row][col] == firstChar && dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
