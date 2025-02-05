package unique_paths

import "fmt"

var seen = map[string]int{}

/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/
func uniquePaths(m int, n int) int {
	// https://lyingheart6174.pixnet.net/blog/post/5122477
	if m == 1 {
		return 1
	}
	if n == 1 {
		return 1
	}
	var key string
	if m >= n {
		key = fmt.Sprintf("%v, %v", m, n)
	} else {
		key = fmt.Sprintf("%v, %v", n, m)
	}
	if _, exists := seen[key]; !exists {
		seen[key] = uniquePaths(m-1, n) + uniquePaths(m, n-1)
	}
	return seen[key]
}
