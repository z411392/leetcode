package maximum_product_subarray

/*
Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.
*/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	currentMax := nums[0]
	currentMin := nums[0]
	maxProduct := nums[0]
	for _, num := range nums[1:] {
		temp := currentMax
		currentMax = max(num, max(temp*num, currentMin*num))
		currentMin = min(num, min(temp*num, currentMin*num))
		maxProduct = max(maxProduct, currentMax)
	}
	return maxProduct
}
