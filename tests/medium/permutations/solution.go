package permutations

/*
Given an array nums of distinct integers, return all the possible
permutations
. You can return the answer in any order.
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	permutations := [][]int{}
	for index, taken := range nums {
		lefted := append([]int{}, nums[:index]...)
		lefted = append(lefted, nums[index+1:]...)
		for _, permutation := range permute(lefted) {
			permutations = append(permutations, append(permutation, taken))
		}
	}
	return permutations
}
