package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	result := make([]int, n)
	prefix := 1
	for i := 0; i < n; i += 1 {
		result[i] = prefix
		prefix *= nums[i]
	}
	suffix := 1
	for i := n - 1; i > -1; i -= 1 {
		result[i] *= suffix
		suffix *= nums[i]
	}
	return result
}
