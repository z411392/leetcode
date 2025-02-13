package count_primes

import "math"

type IsPrime func(n int) bool

func isPrimeFunc() IsPrime {
	primes := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
		23: true,
	}
	return func(n int) bool {
		if isPrime, exists := primes[n]; exists {
			return isPrime
		}
		sqrt := int(math.Floor(math.Sqrt(float64(n))))
		primes[n] = true
		for i := sqrt; i >= 2; i -= 1 {
			if n%i == 0 {
				primes[n] = false
				break
			}
		}
		return primes[n]
	}
}

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

	isPrime := isPrimeFunc()
	primes := 1
	for i := 3; i < n; i += 2 {
		if isPrime(i) {
			primes += 1
		}
	}
	return primes
}
