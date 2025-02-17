package top_k_frequent_elements

func topKFrequent(nums []int, k int) []int {
	// https://leetcode.com/problems/top-k-frequent-elements/submissions/1546110954/
	counters := make(map[int]int)

	for _, num := range nums {
		counters[num]++
	}

	bucket := make([][]int, len(nums))
	for num, counter := range counters {
		bucket[counter-1] = append(bucket[counter-1], num)
	}
	// fmt.Printf("%v\n", bucket)

	taken := make([]int, 0, k)
outer:
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) == 0 {
			continue
		}
		for _, num := range bucket[i] {
			taken = append(taken, num)
			if len(taken) >= k {
				break outer
			}
		}
	}

	return taken
}
