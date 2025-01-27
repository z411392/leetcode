package valid_parentheses

import (
	. "github.com/z411392/leetcode/pkg/data_structure/stack"
)

var (
	pairsMap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
func isValid(s string) bool {
	stack := NewStack[rune]()
	for _, parenthesis := range s {
		expected, isClosing := pairsMap[parenthesis]
		if !isClosing {
			stack.Push(parenthesis)
			continue
		}
		if stack.Size() == 0 {
			return false
		}
		got, _ := stack.Pop()
		if got != expected {
			return false
		}
	}
	return stack.Size() <= 0 // 記得檢查 stack 是不是也一併清空
}
