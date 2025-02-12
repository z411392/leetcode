package factorial_trailing_zeroes

/*
Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.
*/
func trailingZeroes(n int) int {
	// 譬如 n = 30
	// 5 的數目 5(+1), 10(+1), 15(+1), 20(+1), 25(+2), 30(+1) ... 共 7 個
	// 2 的數目 2(+1), 4(+2), 6(+1), 8(+3) ... 30(+1)
	// 幾個零？ 取 2 的數目和 5 的數目間比較小的那個，通常直接取 5 的數目就可以
	zeros := 0
	for {
		if n < 5 {
			break
		}
		n /= 5
		zeros += n
	}
	return zeros
}
