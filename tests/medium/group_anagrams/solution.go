package group_anagrams

/*
Given an array of strings strs, group the
anagrams together. You can return the answer in any order.
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	return [][]string{}
}
