package find_first_and_last_position_of_element_in_sorted_array

/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	left := 0
	right := len(nums) - 1
	for {
		// 如果目標值超出範圍直接返回
		if nums[left] > target {
			return []int{-1, -1}
		}
		if nums[right] < target {
			return []int{-1, -1}
		}
		// 若中心點為目標值則往左右方向擴張
		middle := left + (right-left)/2
		if nums[middle] == target {
			left = middle
			for {
				// fmt.Printf("left=%v\n", left)
				if left-1 < 0 {
					break
				}
				if nums[left-1] != target {
					break
				}
				left -= 1
			}
			right = middle
			for {
				// fmt.Printf("right=%v\n", right)
				if right+1 >= len(nums) {
					break
				}
				if nums[right+1] != target {
					break
				}
				right += 1
			}
			return []int{left, right}
		}
		// 否則不斷限縮範圍直到剩一個點
		// 如果該點不為目標值則返回
		if left == right {
			break
		}
		if nums[middle] > target {
			right = middle - 1
			continue
		}
		if nums[middle] < target {
			left = middle + 1
			continue
		}
	}
	return []int{-1, -1}
}
