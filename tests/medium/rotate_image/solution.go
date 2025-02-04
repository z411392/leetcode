package rotate_image

/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
*/

func rotate(matrix [][]int) {
	// https://bclin.tw/2021/08/24/leetcode-48/
	size := len(matrix)
	if size <= 1 {
		return
	}
	mid := size / 2
	for row := 0; row < size; row += 1 {
		// 將對角線右上的部分換到左下，只換到對角線右邊的部分（對角線是從左上到右下這條）
		for column := row; column < size; column += 1 {
			matrix[row][column], matrix[column][row] = matrix[column][row], matrix[row][column]
		}
		// 左右交換
		for column := 0; column < mid; column++ { // 5/2 = 2 4/2=2 不必包含中心
			mirror := size - 1 - column
			matrix[row][column], matrix[row][mirror] = matrix[row][mirror], matrix[row][column]
		}
	}
}
