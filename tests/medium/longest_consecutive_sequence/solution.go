package longest_consecutive_sequence

/*
Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.
*/
func longestConsecutive(nums []int) int {
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		seen[num] = true
	}
	var longest int
	for num := range seen {
		_, exists := seen[num-1]
		// num-1 存在表不為連續數列的起點
		if exists {
			continue // continue 不會使 n 成爲 n logn 或 n^2
		}
		// 計算連續次數
		length := 1
		for {
			_, exists := seen[num+length]
			if !exists {
				break
			}
			length += 1
		}
		longest = max(longest, length)
	}
	return longest
}
