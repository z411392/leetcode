package string_to_integer_atoi

type ProcessingMode int

const MinUint = uint32(0)
const MaxUint = ^MinUint
const MaxInt = int(int(MaxUint >> 1))
const MinInt = int(MaxInt + 1)

const (
	Whitespace ProcessingMode = iota
	Signedness
	Conversion
	Rounding
)

/*
Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.

The algorithm for myAtoi(string s) is as follows:

1. Whitespace: Ignore any leading whitespace (" ").
2. Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
3. Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
Return the integer as the final result.
*/
func myAtoi(s string) int {
	processingMode := Whitespace
	leadingZero := true
	digits := make([]rune, 0)
	negative := false
	for _, character := range s {
		// 每個 character 會有對應的處理狀態 先前的狀態會持續下去
		if processingMode == Whitespace {
			if character == ' ' {
				// fmt.Printf("character=%c processingMode=%v\n", character, processingMode)
				continue
			}
			processingMode = Signedness
		}
		if processingMode == Signedness {
			if character == '+' {
				// fmt.Printf("character=%c processingMode=%v\n", character, processingMode)
				processingMode = Conversion
				continue
			} else if character == '-' {
				// fmt.Printf("character=%c processingMode=%v\n", character, processingMode)
				processingMode = Conversion
				negative = true
				continue
			} else {
				processingMode = Conversion
			}
		}
		if processingMode == Conversion {
			if character >= '0' && character <= '9' {
				if leadingZero {
					if character == '0' {
						continue
					} else {
						leadingZero = false
					}
				}
				// fmt.Printf("character=%c processingMode=%v\n", character, processingMode)
				digits = append(digits, character)
				continue
			}
			processingMode = Rounding
			break
		}
	}
	// fmt.Printf("negative=%v digits=%v\n", negative, digits)
	n := 0
	for _, digit := range digits {
		n *= 10
		n += int(digit) - int('0')
		// 每次加總完判斷數值大小如果超過範圍就提前跳脫
		if negative {
			if n >= MinInt {
				n = MinInt
				break
			}
		} else {
			if n >= MaxInt {
				n = MaxInt
				break
			}
		}
	}
	if negative {
		return -n
	}
	return n
}
