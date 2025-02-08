package decode_ways

import (
	"strconv"
)

/*
You have intercepted a secret message encoded as a string of numbers. The message is decoded via the following mapping:

"1" -> 'A'

"2" -> 'B'

...

"25" -> 'Y'

"26" -> 'Z'

However, while decoding the message, you realize that there are many different ways you can decode the message because some codes are contained in other codes ("2" and "5" vs "25").

For example, "11106" can be decoded into:

"AAJF" with the grouping (1, 1, 10, 6)
"KJF" with the grouping (11, 10, 6)
The grouping (1, 11, 06) is invalid because "06" is not a valid code (only "6" is valid).
Note: there may be strings that are impossible to decode.

Given a string s containing only digits, return the number of ways to decode it. If the entire string cannot be decoded in any valid way, return 0.

The test cases are generated so that the answer fits in a 32-bit integer.
*/
func numDecodings(s string) int {
	memo := make(map[int]int)
	var countDecodingWays func(s string, index int) int
	countDecodingWays = func(s string, index int) int {
		if val, exists := memo[index]; exists {
			return val
		}
		if index == len(s) {
			return 1
		}
		if s[index] == '0' {
			return 0
		}
		if index == len(s)-1 {
			return 1
		}
		ways := countDecodingWays(s, index+1)
		twoDigits, _ := strconv.Atoi(s[index : index+2])
		if twoDigits <= 26 {
			ways += countDecodingWays(s, index+2)
		}
		memo[index] = ways
		return ways
	}
	return countDecodingWays(s, 0)
}
