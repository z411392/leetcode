package house_robber

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
*/
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	even := 0
	odd := 0
	i := 0
	for {
		if i >= len(nums) {
			break
		}
		even += nums[i]
		if i+1 >= len(nums) {
			break
		}
		odd += nums[i+1]
		i += 2
	}
	return max(even, odd)
}
