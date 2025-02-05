package maximum_subarray

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/
func maxSubArray(nums []int) int {
	// https://www.shubo.io/maximum-subarray-problem-kadane-algorithm/
	result := 0
	ending := 0
	maximum := -10000
	for _, num := range nums {
		ending = max(0, ending+num)
		result = max(result, ending)
		if num > maximum {
			maximum = num
		}
	}
	if result <= 0 {
		return maximum
	}
	return result
}
