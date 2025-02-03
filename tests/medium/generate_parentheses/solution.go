package generate_parentheses

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*/
func generateParenthesis(n int) []string {
	// https://thesumitshrestha.medium.com/day-50-generate-parentheses-81c4454865c2
	permutations := []string{}
	var backtrack func(permutation string, openP int, closeP int)
	backtrack = func(permutation string, openP int, closeP int) {
		reached := openP >= n
		completed := closeP >= openP
		if reached && completed {
			permutations = append(permutations, permutation)
			return
		}
		if !reached {
			backtrack(permutation+"(", openP+1, closeP)
		}
		// 多個情況的遞迴
		if !completed {
			backtrack(permutation+")", openP, closeP+1)
		}
	}
	backtrack("", 0, 0)
	return permutations
}
