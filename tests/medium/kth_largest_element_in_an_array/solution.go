package kth_largest_element_in_an_array

import (
	"math/rand"
	"time"
)

func quickSelect(nums []int, left, right, k int) int {
	pivotIndex := rand.Intn(right-left+1) + left // 隨機選 pivot
	pivot := nums[pivotIndex]

	// 把 pivot 移到最後
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]

	// Partition: 所有比 pivot 大的放左邊
	storeIndex := left
	// storeIndex 表示在整個子陣列內找到幾個比 pivot 大的元素。
	// 若 storeIndex 恰好等於 k - 1，則 pivot 會剛好是第 k 大的元素。
	for i := left; i < right; i++ {
		if nums[i] > pivot { // 找前 k 大，所以比 pivot 大的排左邊
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}

	// 把 pivot 放回正確位置
	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]

	// pivot 剛好是第 k 大？
	if storeIndex == k-1 {
		return storeIndex
	}
	if storeIndex > k-1 {
		return quickSelect(nums, left, storeIndex-1, k) // 往左邊找
	} else {
		return quickSelect(nums, storeIndex+1, right, k) // 往右邊找
	}
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
