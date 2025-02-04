package search_in_rotated_sorted_array

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

func search(nums []int, target int) int {
	// https://www.cnblogs.com/grandyang/p/4325648.html
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := len(nums) - 1
	for {
		if left >= right {
			break
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] > nums[left] {
			if target > nums[left] && target < nums[middle] {
				// 若確定是有序且在範圍內則限縮在此範圍
				right = middle - 1
			} else {
				// 否則隨意嘗試另一邊
				left = middle + 1
			}
		} else {
			if target > nums[middle] && target < nums[right] {
				// 若確定是有序且在範圍內則限縮在此範圍
				left = middle + 1
			} else {
				// 否則隨意嘗試另一邊
				right = middle - 1
			}
		}
	}
	return -1
}
