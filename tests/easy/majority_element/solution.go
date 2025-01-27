package majority_element

func majorityElement(nums []int) int {
	var counters = map[int]int{}
	max := 0
	maxNum := 0
	for _, num := range nums {
		counter := counters[num] + 1
		if counter > max {
			max = counter
			maxNum = num
		}
		counters[num] = counter
	}
	return maxNum
}
