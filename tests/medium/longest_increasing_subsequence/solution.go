package longest_increasing_subsequence

import "sort"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] 表示長度為 i+1 的遞增子序列的最小末尾值
	tails := make([]int, 0)

	for _, num := range nums {
		// 用二分搜尋找到第一個大於等於 num 的位置
		i := sort.SearchInts(tails, num)

		// 如果找到的位置等於 tails 的長度，表示 num 比所有值都大
		if i == len(tails) {
			tails = append(tails, num)
		} else {
			// 更新該位置的值
			tails[i] = num
		}
	}

	return len(tails)
}
