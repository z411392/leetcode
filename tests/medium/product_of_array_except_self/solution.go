package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{0}
	}
	// https://leetcode.com/problems/product-of-array-except-self/solutions/3447796/golang-prefix-sum-easy-to-understand/
	left := make([]int, len(nums))
	left[0] = 1
	right := make([]int, len(nums))
	right[len(nums)-1] = 1

	result := make([]int, len(nums))
	for i := 1; i < len(nums); i += 1 {
		left[i] = left[i-1] * nums[i-1]
		j := len(nums) - i
		right[j-1] = right[j] * nums[j]
	}
	// fmt.Printf("left=%v, right=%v\n", left, right)
	for i := 0; i < len(nums); i += 1 {
		result[i] = left[i] * right[i]
	}
	return result
}
