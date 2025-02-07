package sort_colors

func sortColors(nums []int) {
	for i := 0; i < len(nums)-1; i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] <= nums[j] {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
