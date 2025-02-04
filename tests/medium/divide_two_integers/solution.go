package divide_two_integers

const MinUint = uint32(0)
const MaxUint = ^MinUint
const MaxInt = int(int(MaxUint >> 1))
const MinInt = int(MaxInt + 1)

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.
*/
func divide(dividend int, divisor int) int {
	isDividendNegative := false
	if dividend < 0 {
		isDividendNegative = true
		dividend = -dividend
	}
	isDivisorNegative := false
	if divisor < 0 {
		isDivisorNegative = true
		divisor = -divisor
	}
	quotient := 0
	// 記得處理 divisor = 1 的情形，否則會 timeout
	if divisor == 1 {
		quotient = dividend
	} else {
		// https://haogroot.com/2022/06/18/leetcode-29/
		n := divisor
		m := 1
		divisors := []int{}
		quotients := []int{}
		for {
			if n > dividend {
				break
			}
			divisors = append([]int{n}, divisors...)   // [48, 24, 12, 6, 3]
			quotients = append([]int{m}, quotients...) // [16, 8, 4, 2, 1]
			n <<= 1
			m <<= 1
		}
		// fmt.Printf("%v %v\n", divisors, quotients)
		remained := dividend
		for {
			if remained == 0 {
				break
			}
			found := false
			for index, divisor := range divisors {
				if remained >= divisor {
					found = true
					remained -= divisor
					quotient += quotients[index]
					break
				}
			}
			if !found {
				break
			}
		}
	}
	if isDividendNegative != isDivisorNegative {
		if quotient > MinInt {
			return -MinInt
		}
		return -quotient
	}
	if quotient > MaxInt {
		return MaxInt
	}
	return quotient
}
