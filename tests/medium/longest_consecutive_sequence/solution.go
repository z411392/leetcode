package longest_consecutive_sequence

import (
	"slices"
)

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/
func longestConsecutive(nums []int) int {
	slices.SortFunc(nums, func(a int, b int) int {
		if a < b {
			return -1
		}
		if a == b {
			return 0
		}
		return 1
	})
	var startsWith any = nil
	var lastNum any = nil
	currentLength := 0
	maxLength := 0
	for _, num := range nums {
		_, ok := startsWith.(int)
		// fmt.Printf("數字 %v\n", num)
		if !ok {
			// fmt.Printf("頭一個數字 %v\n", num)
			startsWith = num
			lastNum = num
			currentLength = 1
			maxLength = 1
			continue
		}
		lastNum = lastNum.(int)
		if lastNum == num {
			continue
		}
		if lastNum != num-1 {
			// fmt.Printf("當前數字 %v 與前一個數字 %v 不連續\n", num, lastNum)
			startsWith = num
			lastNum = num
			currentLength = 1
			continue
		}
		currentLength += 1
		if currentLength > maxLength {
			// fmt.Printf("連續次數 %v 大於最大值 %v\n", currentLength, maxLength)
			maxLength = currentLength
		}
		lastNum = num
	}
	return maxLength
}
