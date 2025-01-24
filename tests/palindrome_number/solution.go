package palindrome_number

import (
	"strconv"
)

type Pair struct {
	One     rune
	Another rune
}

func pairsOf(str string) <-chan Pair {
	ch := make(chan Pair)
	right := len(str)
	if right == 0 {
		close(ch)
		return ch
	}
	go func() {
		defer close(ch)
		for left, one := range str {
			right -= 1
			another := rune(str[right])
			ch <- Pair{
				One:     one,
				Another: another,
			}
			if left >= right {
				break
			}
		}
	}()
	return ch
}

/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/
func isPalindrome(x int) bool {
	str := strconv.Itoa(x) // 不能是 str := string(x)
	for pair := range pairsOf(str) {
		if pair.One != pair.Another {
			return false
		}
	}
	return true
}
