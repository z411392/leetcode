package wiggle_sort_ii

import "sort"

func wiggleSort(nums []int) {
	n := len(nums)
	// 複製並排序數組
	temp := make([]int, n)
	copy(temp, nums)
	sort.Ints(temp)

	// 從大到小填充
	// 奇數位置先填充較大的數
	// mid := (n - 1) / 2
	j := n - 1
	for i := 1; i < n; i += 2 {
		nums[i] = temp[j]
		j--
	}
	// 偶數位置填充剩下的數
	for i := 0; i < n; i += 2 {
		nums[i] = temp[j]
		j--
	}
}
