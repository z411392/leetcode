package count_and_say

import (
	"fmt"
	"strings"
)

func rle(str string) string {
	var lastCharacter rune
	repeat := 0
	stringBuilder := &strings.Builder{}
	for _, character := range str {
		if character == lastCharacter {
			repeat += 1
			continue
		}
		if repeat > 0 {
			stringBuilder.WriteString(fmt.Sprintf("%v%c", repeat, lastCharacter))
			repeat = 0
		}
		lastCharacter = character
		repeat = 1
	}
	if repeat > 0 {
		stringBuilder.WriteString(fmt.Sprintf("%v%c", repeat, lastCharacter))
		repeat = 0
	}
	return stringBuilder.String()
}

/*
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the run-length encoding of countAndSay(n - 1).
Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical characters (repeated 2 or more times) with the concatenation of the character and the number marking the count of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23", replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes "23321511".

Given a positive integer n, return the nth element of the count-and-say sequence.
*/
func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	return rle(countAndSay(n - 1))
}
