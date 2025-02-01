package reverse_integer

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
func reverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		// fmt.Printf("digit=%v\n", digit)
		next := result*10 + digit // 合併起來一次檢查
		// fmt.Printf("result=%v next=%v\n", result, next)
		if int32(next)/10 != int32(result) { // 按整數除法運算的特性會忽略餘數 而且此時 result 都還是加上 digit 之前的值
			return 0
		}
		x = x / 10
		result = next
	}
	return result
}
