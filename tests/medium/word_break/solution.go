package word_break

/*
*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/
func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	var fun func(s string) bool
	fun = func(s string) bool {
		if len(s) == 0 {
			return true
		}

		if result, exists := memo[s]; exists {
			return result
		}

		for _, word := range wordDict {
			if len(s) < len(word) {
				continue
			}
			if s[:len(word)] == word {
				if fun(s[len(word):]) {
					memo[s] = true
					return true
				}
			}
		}
		// 一定要記錄失敗的嘗試否則會無限遞迴
		memo[s] = false
		return false
	}
	return fun(s)
}
