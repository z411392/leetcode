package three_sum

import "fmt"

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.
*/
func threeSum(nums []int) [][]int {
	// https://leetcode.com/problems/3sum/solutions/725950/python-5-easy-steps-beats-97-4-annotated/comments/1167162
	nMap := map[int]bool{}
	nSlice := []int{}
	pMap := map[int]bool{}
	pSlice := []int{}
	zeros := 0
	for _, num := range nums {
		if num == 0 {
			zeros += 1
		} else if num < 0 {
			nMap[num] = true
			nSlice = append(nSlice, num)
		} else {
			pMap[num] = true
			pSlice = append(pSlice, num)
		}
	}
	combinations := map[string]bool{}
	results := [][]int{}
	// 雖然順序不重要，但單元測試是簡單用 deep equal，所以還是把區塊的順序調換

	// -p, -q, r
	for i := 0; i < len(nSlice)-1; i += 1 {
		for j := i + 1; j < len(nSlice); j += 1 {
			n := nSlice[i] + nSlice[j]
			_, exists := pMap[-n]
			if !exists {
				continue
			}
			var key string
			if nSlice[i] < nSlice[j] {
				key = fmt.Sprintf("%v %v %v", nSlice[i], nSlice[j], -n)
			} else {
				key = fmt.Sprintf("%v %v %v", nSlice[j], nSlice[i], -n)
			}
			if combinations[key] {
				continue
			}
			combinations[key] = true
			results = append(results, []int{nSlice[i], nSlice[j], -n})
		}
	}
	if zeros > 0 {
		// 0, -n, n
		for num := range nMap {
			_, exists := pMap[-num]
			if exists {
				results = append(results, []int{num, 0, -num})
			}
		}
		// 0, 0, 0
		if zeros > 2 {
			results = append(results, []int{0, 0, 0})
		}
	}

	// p, q, -r
	for i := 0; i < len(pSlice)-1; i += 1 {
		for j := i + 1; j < len(pSlice); j += 1 {
			n := pSlice[i] + pSlice[j]
			_, exists := nMap[-n]
			if !exists {
				continue
			}
			var key string
			if pSlice[i] < pSlice[j] {
				key = fmt.Sprintf("%v %v %v", -n, pSlice[i], pSlice[j])
			} else {
				key = fmt.Sprintf("%v %v %v", -n, pSlice[j], pSlice[i])
			}
			if combinations[key] {
				continue
			}
			combinations[key] = true
			results = append(results, []int{-n, pSlice[i], pSlice[j]})
		}
	}
	return results
}
