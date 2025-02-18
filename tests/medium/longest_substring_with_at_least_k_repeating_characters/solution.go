package longest_substring_with_at_least_k_repeating_characters

func longestSubstring(s string, k int) int {
	// 如果字串長度小於 k，不可能有符合條件的子字串
	if len(s) < k {
		return 0
	}

	// 計算每個字元出現的頻率
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	// 檢查是否所有字元都符合條件
	allValid := true
	for i := 0; i < len(s); i++ {
		if count[s[i]] < k {
			allValid = false
			break
		}
	}

	// 如果所有字元都符合條件，返回整個字串長度
	if allValid {
		return len(s)
	}

	// 找出不符合條件的字元，以此字元分割字串
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {
		if count[s[i]] < k {
			// 遞迴處理子字串
			subLen := longestSubstring(s[start:i], k)
			if subLen > maxLen {
				maxLen = subLen
			}
			start = i + 1
		}
	}

	// 處理最後一個子字串
	if start < len(s) {
		subLen := longestSubstring(s[start:], k)
		if subLen > maxLen {
			maxLen = subLen
		}
	}

	return maxLen
}
