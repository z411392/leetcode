package word_search

func exist(board [][]byte, word string) bool {
	var dfs func(position []int, word string, seen map[[2]int]bool) bool
	dfs = func(position []int, word string, seen map[[2]int]bool) bool {
		if seen[[2]int{position[0], position[1]}] {
			return false
		}
		if position[0] < 0 {
			return false
		}
		if position[0] > len(board)-1 {
			return false
		}
		if position[1] < 0 {
			return false
		}
		if position[1] > len(board[0])-1 {
			return false
		}
		if word[0] != board[position[0]][position[1]] {
			return false
		}
		if len(word) == 1 {
			return true
		}
		seen[[2]int{position[0], position[1]}] = true
		found := dfs([]int{position[0] - 1, position[1]}, word[1:], seen) || dfs([]int{position[0] + 1, position[1]}, word[1:], seen) || dfs([]int{position[0], position[1] - 1}, word[1:], seen) || dfs([]int{position[0], position[1] + 1}, word[1:], seen)
		if found {
			return true
		}
		delete(seen, [2]int{position[0], position[1]})
		return false
	}
	for row, cells := range board {
		for column := range cells {
			found := dfs([]int{row, column}, word, map[[2]int]bool{})
			if found {
				return true
			}
		}
	}
	return false
}
