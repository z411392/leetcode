package sum_of_two_integers

func getSum(a int, b int) int {
	// https://leetcode.com/problems/sum-of-two-integers/solutions/84323/golang-concise-solution-with-an-explanation/
	sum, carry := a^b, a&b
	for {
		if carry == 0 {
			break
		}
		carry <<= 1
		sum, carry = sum^carry, sum&carry
	}
	return sum
}
