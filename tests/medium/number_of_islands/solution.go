package number_of_islands

/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
*/
func numIslands(grid [][]byte) int {
	var dfs func(row, column int)
	dfs = func(row, column int) {
		if row < 0 {
			return
		}
		if column < 0 {
			return
		}
		if row >= len(grid) {
			return
		}
		if column >= len(grid[row]) {
			return
		}
		if grid[row][column] == '0' {
			return
		}
		grid[row][column] = '0'
		dfs(row-1, column)
		dfs(row+1, column)
		dfs(row, column-1)
		dfs(row, column+1)
	}
	nums := 0
	for row, cells := range grid {
		for column, cell := range cells {
			if cell == '0' {
				continue
			}
			nums += 1
			dfs(row, column)
		}
	}
	return nums
}
