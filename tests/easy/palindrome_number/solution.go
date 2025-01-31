package palindrome_number

/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reversed := 0
	original := x
	for {
		reversed *= 10
		reversed += x % 10
		if x < 10 {
			break
		}
		x /= 10
	}
	return original == reversed
}
