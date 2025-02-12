package fraction_to_recurring_decimal

import (
	"strconv"
	"strings"
)

/*
Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return any of them.

It is guaranteed that the length of the answer string is less than 104 for all the given inputs.
*/
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	negative := false
	if numerator < 0 {
		numerator = -numerator
		negative = !negative
	}
	if denominator < 0 {
		denominator = -denominator
		negative = !negative
	}
	digits := &strings.Builder{}
	numeratorMap := map[int]int{}
	decimalIndex := -1
	repeatStart := -1
	repeatEnd := -1
	for {
		// 如果已經處理過相同的被除數表示遇到循環結束
		if value, seen := numeratorMap[numerator]; seen {
			repeatStart = value
			repeatEnd = len(digits.String())
			break
		}
		// 否則假設被除數為循環開始
		numeratorMap[numerator] = len(digits.String())
		// 每次被除數小於除數時補零；第一次補零時保留當時字串的長度（會恰好是小數點的 index）
		if numerator < denominator {
			digits.WriteRune('0')
		} else {
			remainder := numerator % denominator
			quotient := numerator / denominator
			// fmt.Printf("%v/%v=%v\n", numerator, denominator, quotient)
			// 餘數為零時表示能整除不必考慮無限循環的問題
			if remainder == 0 {
				digits.WriteString(strconv.Itoa(quotient))
				break
			}
			numerator = remainder
			digits.WriteString(strconv.Itoa(quotient))
		}
		if decimalIndex == -1 {
			decimalIndex = len(digits.String())
		}
		numerator *= 10
	}
	result := digits.String()
	// fmt.Printf("result=%s, decimalIndex=%v, repeatStart=%v, repeatEnd=%v, negative=%v\n", result, decimalIndex, repeatStart, repeatEnd, negative)
	// 如果從整數部分就已經開始循環，要將循環的部分往右移
	if repeatStart != -1 && decimalIndex != -1 && repeatStart < decimalIndex {
		shift := decimalIndex - repeatStart
		result += result[repeatStart:decimalIndex]
		repeatStart += shift
		repeatEnd += shift
	}
	if repeatEnd != -1 {
		result = result[:repeatEnd] + ")"
	}
	if repeatStart != -1 {
		result = result[:repeatStart] + "(" + result[repeatStart:]
	}
	if decimalIndex != -1 {
		result = result[:decimalIndex] + "." + result[decimalIndex:]
	}
	if negative {
		result = "-" + result
	}
	return result
}
