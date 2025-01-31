package longest_palindromic_substring

import (
	"strings"
)

func joinWithSeparator(str string, separator string) string {
	stringBuilder := &strings.Builder{}
	stringBuilder.WriteString(separator)
	stringBuilder.WriteString(strings.Join(strings.Split(str, ""), separator))
	stringBuilder.WriteString(separator)
	return stringBuilder.String()
}

/*
Given a string s, return the longest palindromic substring in s.
*/
func longestPalindrome(inputString string) string {
	// https://www.youtube.com/watch?v=SHrjQKUNl1g
	processedString := joinWithSeparator(inputString, "#")
	palindromeRadiuses := make([]int, len(processedString))
	center := 0
	rightBoundary := 0
	maxPalindromeLength := 0
	palindrome := ""
	for index := range len(processedString) {
		// fmt.Printf("index=%v, rightBoundary=%v\n", index, rightBoundary)
		if index <= rightBoundary {
			// fmt.Printf("palindromeRadiuses[index]=%v\n", palindromeRadiuses[index])
			palindromeRadiuses[index] = palindromeRadiuses[2*center-index] // 假設 center 已更新，這邊 index 的永遠是傘蓋的右邊；用傘蓋左邊的值來作預設值
			if rightBoundary-index < palindromeRadiuses[index] {           // 如果在散蓋左側的這朵蘑菇只是部分覆蓋，只能先假定蘑菇半徑至少有到最大蘑菇的邊界
				palindromeRadiuses[index] = rightBoundary - index
			}
		}
		for {
			left := index - palindromeRadiuses[index] - 1
			right := index + palindromeRadiuses[index] + 1
			// fmt.Printf("left=%v, right=%v\n", left, right)
			if left < 0 {
				break
			}
			if right >= len(processedString) {
				break
			}
			if processedString[left] != processedString[right] {
				break
			}
			palindromeRadiuses[index] += 1
			// fmt.Printf("palindromeRadiuses[index]=%v\n", palindromeRadiuses[index])
		}
		// fmt.Printf("index+palindromeRadiuses[index]=%v rightBoundary=%v\n", index+palindromeRadiuses[index], rightBoundary)
		if palindromeRadiuses[index] > maxPalindromeLength {
			maxPalindromeLength = palindromeRadiuses[index]
			palindrome = processedString[index-palindromeRadiuses[index] : index+palindromeRadiuses[index]+1]
		}
		if index+palindromeRadiuses[index] > rightBoundary {
			rightBoundary = index + palindromeRadiuses[index]
			center = index
		}
	}
	// fmt.Printf("%v\n", palindromeRadiuses)
	return strings.Replace(palindrome, "#", "", -1)
}
