package next_permutation

import (
	"fmt"
	"reflect"
	"slices"
)

func permutationsFunc(nums []int) func() [][]int {
	length := len(nums)
	seen := map[string]bool{}
	copied := make([]int, len(nums))
	copy(copied, nums)
	slices.Sort(copied)
	var permutations func(nums []int) [][]int
	permutations = func(nums []int) [][]int {
		result := [][]int{}
		if len(nums) == 1 {
			permutation := []int{nums[0]}
			result = append(result, permutation)
			return result
		}
		for index, num := range nums {
			sub := slices.Concat(nums[:index], nums[index+1:])
			for _, permutation := range permutations(sub) {
				permutation = slices.Concat([]int{num}, permutation)
				if length != len(permutation) {
					result = append(result, permutation)
					continue
				}
				signature := fmt.Sprint(permutation)
				if _, exists := seen[signature]; exists {
					continue
				}
				seen[signature] = true
				result = append(result, permutation)
			}
		}
		return result
	}
	return func() [][]int {
		return permutations(copied)
	}
}

func nextPermutation(nums []int) {
	permutations := permutationsFunc(nums)
	all := permutations()
	found := -1
	for index, permutation := range all {
		// fmt.Printf("%v\n", permutation)
		if reflect.DeepEqual(nums, permutation) {
			found = index
			break
		}
	}
	next := (found + 1) % len(all)
	copy(nums, all[next])
}
