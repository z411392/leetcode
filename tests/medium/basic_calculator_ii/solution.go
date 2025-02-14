package basic_calculator_ii

import (
	"unicode"
)

func calculate(s string) int {

	number := 0
	sign := '+'
	change := 0
	result := 0

	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if c == ' ' {
			continue
		}
		if unicode.IsDigit(c) {
			number = number*10 + int(c-'0')
			continue
		}
		// fmt.Printf("c=%c sign=%c\n", c, sign)

		switch sign {
		case '+':
			// 遇到 + - 時將鏈鎖改動的 change 加回 result 上
			result += change
			// fmt.Printf("result+=%v\n", change)
			// 並初始化下一次的 changing（如果是負數，則前面掛負號）
			change = number
			// fmt.Printf("result+=%v, change=%v\n", result, change)
		case '-':
			result += change
			// fmt.Printf("result+=%v\n", change)
			change = -number
			// fmt.Printf("result=%v, change=%v\n", result, change)
		case '*':
			// 遇到 * / 時不結算 result，而是鏈鎖改變 change
			change = change * number
			// fmt.Printf("change=%v\n", change)
		case '/':
			change = change / number
			// fmt.Printf("change=%v\n", change)
		}
		sign = c
		number = 0
		// fmt.Printf("sign=%c, number=%v\n", sign, number)

	}
	switch sign {
	case '+':
		result += change
		change = number
	case '-':
		result += change
		change = -number

	case '*':
		change = change * number
	case '/':
		change = change / number
	}

	return result + change
}
