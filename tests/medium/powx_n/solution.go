package count_and_say

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	negative := false
	if x < 0 {
		if n&1 == 1 {
			negative = true
		}
		x = -x
	}
	result := float64(1)
	if x != 1 {
		if n < 0 {
			n = -n
			x = 1 / x
		}
		// https://leetcode.com/problems/powx-n/solutions/1337794/java-c-simple-o-log-n-easy-faster-than-100-explained/
		for {
			if n <= 0 {
				break
			}
			if n&1 == 1 {
				result = result * x
			}
			x *= x
			n >>= 1
		}
	}
	if negative {
		return -result
	}
	return result
}
