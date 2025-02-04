package permutations

/*
Given an array nums of distinct integers, return all the possible
permutations
. You can return the answer in any order.
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	return [][]int{}
}
