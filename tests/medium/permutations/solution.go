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
		lefted := append(append([]int{}, nums[:index]...), nums[index+1:]...)
		for _, permutation := range permute(lefted) {
			permutation = append([]int{taken}, permutation...)
			// fmt.Printf("taken=%v lefted=%v permutation=%v\n", taken, lefted, permutation)
			permutations = append(permutations, permutation)
		}
	}
	return permutations
}
