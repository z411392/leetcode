package merge_intervals

import (
	"slices"
)

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/
func merge(intervals [][]int) [][]int {
	// https://leetcode.com/problems/merge-intervals/solutions/810795/python-js-go-c-by-sort-merge-w-visualization/

	// 不先排序的話 merged 結果之間要再做一次合併
	slices.SortFunc(intervals, func(one, another []int) int {
		if one[0] < another[0] {
			return -1
		}
		if one[0] > another[0] {
			return 1
		}
		if one[1] < another[1] {
			return -1
		}
		if one[1] > another[1] {
			return 1
		}
		return 0
	})
	index := 0
	merged := [][]int{intervals[0]}
	for _, newOne := range intervals[1:] {
		// 減少 assignment 時對陣列的複製可以加快速度
		if newOne[0] <= merged[index][1] {
			// fmt.Printf("重疊\n")
			if newOne[1] > merged[index][1] { // newOne 可能比較短
				merged[index][1] = newOne[1]
			}
			continue
		}
		// fmt.Printf("不重疊\n")
		merged = append(merged, newOne)
		index += 1
	}

	return merged
}
