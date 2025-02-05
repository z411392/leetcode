package count_and_say

import "math"

// https://stackoverflow.com/questions/47969385/go-float-comparison
const epsilon = 1e-9

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	result := float64(x)
	if x != 1 {
		dividing := false
		if n < 0 {
			dividing = true
			n = -n
		}
		if dividing {
			for range n + 1 {
				result = result / x
				if result <= epsilon {
					result = 0
					break
				}
			}
		} else {
			for range n - 1 {
				result = result * x
				if result >= math.MaxFloat64 {
					result = math.MaxFloat64
					break
				}
			}
		}
	}
	if n%2 == 1 && negative {
		result = -result
	}
	return result
}
