package count_and_say

import (
	"math/big"
)

/*
Implement pow(x, n), which calculates x raised to the power n (i.e., xn).
*/
func myPow(x float64, n int) float64 {
	value := big.NewFloat(x)
	if n == 0 {
		return 1
	}
	nagative := false
	if n < 0 {
		nagative = true
		n = -n
	}
	result := big.NewFloat(x)
	if nagative {
		for range n + 1 {
			// fmt.Printf("before=%v\n", result)
			result = big.NewFloat(0).Quo(result, value)
			// fmt.Printf("after=%v\n", result)
		}
	} else {
		for range n - 1 {
			// fmt.Printf("before=%v\n", result)
			result = big.NewFloat(0).Mul(result, value)
			// fmt.Printf("after=%v\n", result)
		}
	}
	y, _ := result.Float64()
	return y
}
