package roman_to_integer

var (
	romanToIntMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
*/
func romanToInt(s string) int {
	represent := 0
	previous := 0
	// 前面的字母如果右邊比較大要扣掉左邊（否則加上左邊）
	for _, character := range s {
		current := romanToIntMap[string(character)]
		if previous != 0 {
			if current > previous {
				represent -= previous
			} else {
				represent += previous
			}
		}
		previous = current
	}
	// 最後一個字母永遠是加上去的
	endCharacter := rune(s[len(s)-1])
	represent += romanToIntMap[string(endCharacter)]
	return represent
}
