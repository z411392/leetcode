package reverse_integer

import (
	"strconv"
	"strings"
)

/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/
func reverse(x int) int {
	if x == 0 {
		return 0
	}
	str := strconv.Itoa(x)
	stringBuilder := &strings.Builder{}
	if strings.HasPrefix(str, "-") {
		stringBuilder.WriteRune('-')
		str = strings.ReplaceAll(str, "-", "")
	}
	// fmt.Printf("%v %v\n", x, str)
	for i := len(str) - 1; i > -1; i-- {
		stringBuilder.WriteByte(str[i])
	}
	stringified := stringBuilder.String()
	result, err := strconv.ParseInt(stringified, 10, 32)
	if err != nil {
		return 0
	}
	return int(result)
}
