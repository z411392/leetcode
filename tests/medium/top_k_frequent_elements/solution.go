package top_k_frequent_elements

import (
	"maps"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	// https://leetcode.com/problems/top-k-frequent-elements/solutions/519149/go-8ms-using-priorityqueue/
	counters := make(map[int]int)
	for _, num := range nums {
		counters[num] += 1
	}
	keys := slices.Collect(maps.Keys(counters))
	slices.SortFunc(keys, func(a, b int) int {
		cmp := counters[b] - counters[a]
		if cmp == 0 {
			return a - b
		}
		return cmp
	})
	return keys[:k]
}
