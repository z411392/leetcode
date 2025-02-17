package wiggle_sort_ii

import (
	"slices"
)

func wiggleSort(nums []int) {
	cloned := slices.Clone(nums)
	slices.Sort(cloned)
	left, right := cloned[:len(nums)/2], cloned[len(nums)/2:]
	i := -1
	for {
		if i == len(nums)-1 {
			break
		}
		if len(left) > 0 {
			i += 1
			nums[i] = left[0]
			left = left[1:]
		}
		if len(right) > 0 {
			i += 1
			nums[i] = right[0]
			right = right[1:]
		}
	}
}
