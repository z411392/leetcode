package kth_largest_element_in_an_array

import (
	"math/rand"
	"time"
)

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return left
	}

	pivotIndex := rand.Intn(right-left+1) + left
	pivot := nums[pivotIndex]

	// 三路切分: lt 表示小於 pivot 的右邊界
	//          gt 表示大於 pivot 的左邊界
	lt, i, gt := left, left, right
	for i <= gt {
		if nums[i] > pivot {
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		} else if nums[i] < pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		} else {
			i++
		}
	}
	if k-1 < lt {
		return quickSelect(nums, left, lt-1, k)
	}
	if k-1 > gt {
		return quickSelect(nums, gt+1, right, k)
	}
	return lt
}

/*
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?
*/
func findKthLargest(nums []int, k int) int {
	rand.New(rand.NewSource(time.Now().UnixNano())) // https://gabrieleromanato.name/go-rand-seed-has-been-deprecated-a-solution
	index := quickSelect(nums, 0, len(nums)-1, k)
	return nums[index]
}
