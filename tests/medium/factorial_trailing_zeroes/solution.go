package factorial_trailing_zeroes

/*
Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.
*/
func trailingZeroes(n int) int {
	return min(n/5, n/2)
}
