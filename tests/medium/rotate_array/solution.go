package rotate_array

import "slices"

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
func rotate(nums []int, k int) {
	shift := k % len(nums)
	for i, num := range slices.Concat(nums[len(nums)-shift:], nums[:len(nums)-shift]) {
		nums[i] = num
	}
}
