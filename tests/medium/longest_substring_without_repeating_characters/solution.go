package longest_substring_without_repeating_characters

//lint:file-ignore ST1001 _

/*
Given a string s, find the length of the longest substring without repeating characters.
*/
func lengthOfLongestSubstring(s string) int {
	// 為每個 index, character 維護各自的 map，一但出現重複的字母就不繼續滾 character
	maxLength := 0
	sets := make(map[int](map[rune]bool))
	closed := make(map[int]bool)
	for i, c := range s {
		for j, set := range sets {
			if closed[j] {
				continue
			}
			if _, ok := set[c]; ok {
				closed[j] = true
				continue
			}
			set[c] = true
			if len(set) > maxLength {
				// for key := range set {
				// 	fmt.Printf("%v", string(key))
				// }
				// fmt.Printf("\n")
				maxLength = len(set)
			}
		}
		sets[i] = make(map[rune]bool)
		sets[i][c] = true
		if maxLength == 0 {
			maxLength = 1
		}
	}
	return maxLength
}
