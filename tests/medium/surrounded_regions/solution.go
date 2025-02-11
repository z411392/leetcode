package surrounded_regions

func dfs(board [][]byte, survived [][]int) (fun func(i, j int)) {
	fun = func(i, j int) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
			return
		}
		if board[i][j] != 'O' {
			return
		}
		if survived[i][j] == 1 {
			return
		}
		survived[i][j] = 1
		// fmt.Printf("(%v, %v)=1\n", i, j)
		fun(i-1, j)
		fun(i+1, j)
		fun(i, j-1)
		fun(i, j+1)
	}
	return fun
}

/*
You are given an m x n matrix board containing letters 'X' and 'O', capture regions that are surrounded:

Connect: A cell is connected to adjacent cells horizontally or vertically.
Region: To form a region connect every 'O' cell.
Surround: The region is surrounded with 'X' cells if you can connect the region with 'X' cells and none of the region cells are on the edge of the board.
To capture a surrounded region, replace all 'O's with 'X's in-place within the original board. You do not need to return anything.
*/
func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	survived := [][]int{}
	for i := 0; i < len(board); i += 1 {
		row := []int{}
		for j := 0; j < len(board[i]); j += 1 {
			row = append(row, 0)
		}
		survived = append(survived, row)
	}
	fun := dfs(board, survived)
	for i := 0; i < len(board); i += 1 {
		for j := 0; j < len(board[i]); j += 1 {
			if !(i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1) {
				continue
			}
			value := board[i][j]
			// 題目只考慮 O 的存活，一個 X 被 O 包圍並不會變成 O
			if value != 'O' {
				continue
			}
			// 從邊界開始對所有 O 作 dfs 確認最終存活的 O 有哪些其餘的全打 X 就可以
			fun(i, j)
		}
	}
	for i := 0; i < len(board); i += 1 {
		for j := 0; j < len(board[i]); j += 1 {
			if survived[i][j] == 1 {
				continue
			}
			board[i][j] = 'X'
		}
	}
}
