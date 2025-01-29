package longest_substring_without_repeating_characters

//lint:file-ignore ST1001 _

/*
Given a string s, find the length of the longest substring without repeating characters.
*/
func lengthOfLongestSubstring(s string) int {
	// 維護 start, end 兩個 cursors
	start := -1
	end := 0
	// 每當有重複字母出現時 start cursor 就往右移 最大字母長度會重新開始計算 不過先前已經累計的仍會保存
	seen := make(map[rune]int)
	maxLength := 0
	for {
		// end cursor 觸底時結束
		if end >= len(s) {
			break
		}
		c := rune(s[end])
		index, exists := seen[c]
		if exists && index > start {
			start = index
			continue
		}
		maxLength = max(maxLength, end-start)
		seen[c] = end
		end += 1
	}
	return maxLength
}
