package rotate_array

import "slices"

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
func rotate(nums []int, k int) {
	shift := k % len(nums)
	before := make([]int, shift)
	copy(before, nums[len(nums)-shift:])
	after := make([]int, len(nums[:len(nums)-shift]))
	copy(after, nums[:len(nums)-shift])
	for i, num := range slices.Concat(before, after) {
		nums[i] = num
	}
}
