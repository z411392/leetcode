package subsets

func subsets(nums []int) [][]int {
	subsets := [][]int{{}}
	if len(nums) == 0 {
		return subsets
	}
	for _, num := range nums {
		for j := range subsets {
			subset := append([]int{}, subsets[j]...)
			// 保留原有的 subsets 直接再複製一份
			subsets = append(subsets, append(subset, num))
		}
	}
	return subsets
}
