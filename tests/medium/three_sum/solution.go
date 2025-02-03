package three_sum

import (
	"fmt"
	"sort"
)

func moveLeft(nums []int, left int, right int) int {
	held := nums[left]
	for {
		left += 1
		// fmt.Printf("left=%v\n", left)
		if left >= right {
			break
		}
		current := nums[left]
		if current != held {
			break
		}
	}
	return left
}

func moveRight(nums []int, left int, right int) int {
	held := nums[right]
	for {
		right -= 1
		// fmt.Printf("right=%v\n", right)
		if left >= right {
			break
		}
		current := nums[right]
		if current != held {
			break
		}
	}
	return right
}

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(nums []int) [][]int {
	// https://leetcode.com/problems/3sum/solutions/295224/go-768-ms-faster-than-100-00
	results := make([][]int, 0)
	seen := map[string]bool{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i += 1 {
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1
		for {
			if left >= right {
				break
			}
			found := nums[left] + nums[right]
			if found == target {
				key := fmt.Sprintf("%v %v %v", nums[i], nums[left], nums[right])
				if !seen[key] {
					results = append(results, []int{nums[i], nums[left], nums[right]})
					seen[key] = true
				}
				left = moveLeft(nums, left, right)
				right = moveRight(nums, left, right)
			} else if found > target {
				right = moveRight(nums, left, right)
			} else {
				left = moveLeft(nums, left, right)
			}
		}
	}
	return results
}
