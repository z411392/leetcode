package binary_search

/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	index := len(nums) / 2
	if nums[index] == target {
		return index
	}
	if index == 0 {
		return -1
	}
	if target < nums[index] {
		return search(nums[:index], target)
	} else {
		subindex := search(nums[index:], target) // 若查無直接回傳 -1 無須校正
		if subindex == -1 {
			return -1
		}
		return index + subindex // 右半邊要加上基底偏移
	}
}
