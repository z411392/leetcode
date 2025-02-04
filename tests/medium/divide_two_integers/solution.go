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
		for {
			// fmt.Printf("%v %v\n", dividend, quotient)
			if dividend < divisor {
				break
			}
			dividend -= divisor
			quotient += 1
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
