package game_of_life

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	populationAt := func(row, column int) int {
		population := 0
		dirs := [][]int{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			// {0, 0},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		}
		for _, dir := range dirs {
			neighborRow := row + dir[0]
			if neighborRow < 0 || neighborRow >= m {
				continue
			}
			neighborColumn := column + dir[1]
			if neighborColumn < 0 || neighborColumn >= n {
				continue
			}
			population += board[neighborRow][neighborColumn]
		}
		return population
	}
	changing := make([][]int, m)
	for row, cells := range board {
		cellsChanging := make([]int, n)
		for column, cell := range cells {
			if cell == 0 {
				population := populationAt(row, column)
				// fmt.Printf("(%v, %v)=%v\n", row, column, population)
				if population == 3 {
					cellsChanging[column] = 1
					continue
				}
				cellsChanging[column] = 0
				continue
			}
			if cell == 1 {
				population := populationAt(row, column)
				// fmt.Printf("(%v, %v)=%v\n", row, column, population)
				if population < 2 {
					cellsChanging[column] = 1
					continue
				}
				if population > 3 {
					cellsChanging[column] = 1
					continue
				}
				cellsChanging[column] = 0
				continue
			}
		}
		// fmt.Printf("%v\n", cellsChanging)
		changing[row] = cellsChanging
	}
	for row, cells := range changing {
		for column, cell := range cells {
			if cell == 0 {
				continue
			}
			if board[row][column] == 0 {
				board[row][column] = 1
				continue
			}
			if board[row][column] == 1 {
				board[row][column] = 0
				continue
			}
		}
	}
}
