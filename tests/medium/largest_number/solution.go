package largest_number

import (
	"slices"
	"strconv"
	"strings"
)

/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.
*/
func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return "0"
	}
	numbers := []string{}
	otherThanZero := false
	for _, num := range nums {
		if num != 0 {
			otherThanZero = true
		}
		number := strconv.Itoa(num)
		numbers = append(numbers, number)
	}
	// 沒有不為零的數字時，直接回傳零（否則會得到 00，不是正確的數值）
	if !otherThanZero {
		return "0"
	}
	slices.SortFunc(numbers, func(a, b string) int {
		n1, _ := strconv.Atoi(a + b)
		n2, _ := strconv.Atoi(b + a)
		return -(n1 - n2) // 如果 n1 > n2 時我們希望 n1 在左邊(-1)
	})
	result := &strings.Builder{}
	for _, number := range numbers {
		result.WriteString(number)
	}
	return result.String()
}
