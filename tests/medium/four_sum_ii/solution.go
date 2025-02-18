package four_sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// https://leetcode.com/problems/4sum-ii/solutions/5220463/go-1-21-clean-code/
	data := make(map[int]int, len(nums1)*len(nums2)) // 最多只寫入 nxn 個值 有先宣告好的話可以省時間
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			// fmt.Printf("%v\n", num1+num2)
			data[num1+num2]++
		}
	}
	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			// fmt.Printf("%v\n", -num3-num4)
			count += data[-num3-num4]
		}
	}
	return count
}
