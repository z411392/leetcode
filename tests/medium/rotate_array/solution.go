package rotate_array

/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/
func rotate(nums []int, k int) {
	original := make([]int, len(nums))
	copy(original, nums)
	shift := k % len(nums)
	for i, num := range original {
		j := (i + shift) % len(nums)
		nums[j] = num
	}
}
