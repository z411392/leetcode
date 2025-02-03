package letter_combinations_of_a_phone_number

var keyboard = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"n", "m", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func combinationsOf(digits string) []string {
	if digits == "" {
		return []string{}
	}
	characters := keyboard[rune(digits[0])]
	if len(digits) == 1 {
		return characters
	}
	combinations := []string{}
	for _, character := range characters {
		for _, combination := range combinationsOf(digits[1:]) {
			combinations = append(combinations, character+combination)
		}
	}
	return combinations
}

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
*/
func letterCombinations(digits string) []string {
	return combinationsOf(digits)
}
