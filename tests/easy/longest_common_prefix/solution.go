package longest_common_prefix

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
func longestCommonPrefix(strs []string) string {
	index := 0
	commonPrefix := ""
	// 先做一層 for 迴圈一個一個比較字母（不特別先比較／儲存最短字串長度）
	for {
		comparedTo := ""
		for _, str := range strs {
			// 如果已經有字串提前比完了就不會有更長的 common prefix
			if index > len(str)-1 {
				return commonPrefix
			}
			character := string(rune(str[index]))
			// 按第一個拿到的字串中的字母當作比較基礎
			if comparedTo == "" {
				comparedTo = character
				continue
			}
			// 如果已經發現有字母不一樣就不會有更長的 common prefix
			if character != comparedTo {
				return commonPrefix
			}
		}
		commonPrefix += comparedTo
		index += 1
	}
}
