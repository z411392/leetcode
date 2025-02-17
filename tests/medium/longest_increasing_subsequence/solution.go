package longest_increasing_subsequence

import (
	"slices"
)

func lengthOfLIS(nums []int) int {
	// https://medium.com/%E6%8A%80%E8%A1%93%E7%AD%86%E8%A8%98/leetcode-%E8%A7%A3%E9%A1%8C%E7%B4%80%E9%8C%84-300-longest-increasing-subsequence-f160358db4d1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i += 1 {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i += 1 {
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}
			// fmt.Printf("nums[%v]=%v, nums[%v]=%v\n", i, nums[i], j, nums[j])
			dp[i] = max(dp[i], dp[j]+1)
		}
	}
	// fmt.Printf("%v\n", dp)
	return slices.Max(dp)
}
