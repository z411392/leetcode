package count_primes

import (
	"math"
)

/*
Given an integer n, return the number of prime numbers that are strictly less than n.
*/
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	if n == 3 {
		return 1
	}
	isNotPrime := make([]bool, n/2) // 譬如 n = 5, 應該只有 1, 3 兩個奇數要檢查 (5 不被包含)
	sqrt := int(math.Floor(math.Sqrt(float64(n))))
	for num := 3; num <= sqrt; num += 2 {
		i := num / 2
		if isNotPrime[i] {
			// 因為已經被更細粒度的處理過，可以直接忽略
			continue
		}
		for padding := num * num; padding < n; padding += 2 * num { // 略過 num 自身 且 num*num 必不可能被其他質數處理過 因為不可能為偶數 每次直接加 2 * num
			i := padding / 2
			if isNotPrime[i] {
				continue
			}
			isNotPrime[i] = true
		}
	}
	count := 0
	for _, notPrime := range isNotPrime {
		if notPrime {
			continue
		}
		count += 1
	}
	return count
}
