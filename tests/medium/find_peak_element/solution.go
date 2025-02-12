package find_peak_element

/*
A peak element is an element that is strictly greater than its neighbors.

Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

You must write an algorithm that runs in O(log n) time.
*/
func findPeakElement(nums []int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] >= nums[1] {
			return 0
		}
		return 1
	}
	// 依據題意，資料被賦予必定存在峰值的前提（會有波，可能一個到多個）。
	// 可以每次用中點分兩邊，若左邊比右邊大，則比較大的波（或者只有一個波時的最大值）在左邊，反之在右邊。如是反復。
	mid := len(nums) / 2
	if nums[mid-1] > nums[mid] {
		return findPeakElement(nums[:mid])
	} else if nums[mid+1] > nums[mid] {
		return mid + 1 + findPeakElement(nums[mid+1:])
	}
	return mid
}
