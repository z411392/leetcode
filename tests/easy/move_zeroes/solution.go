package move_zeroes

func swap(nums []int, i int) {
	for j := i; j < len(nums); j += 1 {
		if nums[j] == 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		break
	}
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i += 1 {
		if nums[i] != 0 {
			continue
		}
		swap(nums, i)
	}
}
