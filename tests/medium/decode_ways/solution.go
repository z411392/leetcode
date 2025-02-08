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
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	// fmt.Printf("%s\n", s)
	if len(s) == 1 {
		return 1
	}
	ways := 0
	n, _ := strconv.Atoi(s[:2])
	if n <= 26 {
		if len(s[2:]) == 0 {
			ways += 1
		} else {
			ways += numDecodings(s[2:])
		}
	}
	ways += numDecodings(s[1:])
	return ways
}
