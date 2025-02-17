package increasing_triplet_subsequence

func increasingTriplet(nums []int) bool {
	for i := len(nums) - 1; i >= 2; i -= 1 {
		for j := len(nums[:i]) - 1; j >= 1; j -= 1 {
			if nums[i] <= nums[j] {
				continue
			}
			for k := len(nums[:j]) - 1; k >= 0; k -= 1 {
				if nums[j] <= nums[k] {
					continue
				}
				// fmt.Printf("nums[%v]=%v, nums[%v]=%v, nums[%v]=%v\n", i, nums[i], j, nums[j], k, nums[k])
				return true
			}
		}
	}
	return false
}
