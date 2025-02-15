package perfect_squares

/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.
*/
func numSquares(n int) int {
	// https://mikalibrary.com/posts/leetcode/279/
	// https://medium.com/ho-japan/leetcode%E9%A1%8C%E8%A7%A3-2-%E5%8B%95%E6%85%8B%E8%A6%8F%E5%8A%83-dynamic-programming-279-perfect-squares-b257be1170cc
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		minimun := i
		j := 1
		for {
			k := i - j*j
			if k < 0 { // k 可以等於 0，表示當前的 i 是某個 j 的完全平方數；而因為 dp[0]=0，完全平方數只要用 dp[0]+1 個，符合定義。
				break
			}
			value := dp[k] + 1
			if value < minimun {
				minimun = value
			}
			// fmt.Printf("i=%v, j=%v, k=%v, value=%v, minimun=%v\n", i, j, k, value, minimun)
			j += 1
		}
		dp[i] = minimun
	}
	return dp[n]
}
