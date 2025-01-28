package maximum_number_of_fish_in_a_grid

// https://github.com/doocs/leetcode/blob/main/solution/2600-2699/2658.Maximum%20Number%20of%20Fish%20in%20a%20Grid/README.md

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上、下、左、右

// func dfs(grid [][]int, visited [][]bool, row, column, null int) int {
// 	rows := len(grid)
// 	columns := len(grid[0])
// 	if row < 0 {
// 		return 0
// 	}
// 	if row >= rows {
// 		return 0
// 	}
// 	if column < 0 {
// 		return 0
// 	}
// 	if column >= columns {
// 		return 0
// 	}
// 	if grid[row][column] == null { // 若為 0 則不加入 quueue 所以由 0 再往旁邊擴散
// 		return 0
// 	}
// 	if visited[row][column] {
// 		return 0
// 	}
// 	visited[row][column] = true
// 	sum := grid[row][column]
// 	for _, dir := range directions {
// 		nextRow, nextColumn := row+dir[0], column+dir[1]
// 		sum += dfs(grid, visited, nextRow, nextColumn, null)
// 	}
// 	return sum
// }

func bfs(grid [][]int, visited [][]bool, row, column, null int) int {
	rows := len(grid)
	columns := len(grid[0])
	queue := [][]int{{row, column}}
	visited[row][column] = true
	sum := 0
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		row, column := cell[0], cell[1]
		sum += grid[row][column]
		for _, dir := range directions {
			nextRow, nextColumn := row+dir[0], column+dir[1]
			if nextRow < 0 {
				continue
			}
			if nextRow >= rows {
				continue
			}
			if nextColumn < 0 {
				continue
			}
			if nextColumn >= columns {
				continue
			}
			if grid[nextRow][nextColumn] == null { // 若為 0 則不加入 quueue 所以由 0 再往旁邊擴散
				continue
			}
			if visited[nextRow][nextColumn] {
				continue
			}
			visited[nextRow][nextColumn] = true
			queue = append(queue, []int{nextRow, nextColumn}) // 加入 queue 視為繼續 bfs
		}
	}
	return sum
}

const null = 0

/*
You are given a 0-indexed 2D matrix grid of size m x n, where (r, c) represents:

A land cell if grid[r][c] = 0, or
A water cell containing grid[r][c] fish, if grid[r][c] > 0.
A fisher can start at any water cell (r, c) and can do the following operations any number of times:

Catch all the fish at cell (r, c), or
Move to any adjacent water cell.
Return the maximum number of fish the fisher can catch if he chooses his starting cell optimally, or 0 if no water cell exists.

An adjacent cell of the cell (r, c), is one of the cells (r, c + 1), (r, c - 1), (r + 1, c) or (r - 1, c) if it exists.
*/
func findMaxFish(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for row := range visited {
		visited[row] = make([]bool, columns)
	}
	max := 0
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			// 以不為 0 的點當作起點
			if grid[row][column] == 0 {
				continue
			}
			// 且不能已經被處理過
			if visited[row][column] {
				continue
			}
			current := bfs(grid, visited, row, column, null)
			// fmt.Printf("grid[%v][%v]=%v\n", row, column, current)
			if current > max {
				max = current
			}
		}
	}
	return max
}
