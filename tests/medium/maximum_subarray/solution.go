package maximum_subarray

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.
*/
func maxSubArray(nums []int) int {
	// https://www.shubo.io/maximum-subarray-problem-kadane-algorithm/
	// https://leetcode.com/problems/maximum-subarray/description/
	acc := nums[0]
	sum := 0
	for _, num := range nums {
		if sum < 0 {
			sum = 0
		}
		// sum <=0 && num < 0 時也不能略過，因為如果全都是負數還是得找出一個值比較小的負數
		sum += num
		acc = max(acc, sum)
	}
	return acc
}
