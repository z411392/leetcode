package palindrome_partitioning

type IsPalindrome = func(s string) bool

func isPalindromeFunc() IsPalindrome {
	cache := map[string]bool{}
	return func(s string) bool {
		if len(s) <= 1 {
			return true
		}
		if len(s) == 2 {
			return s[0] == s[1]
		}
		if isPalindrome, exists := cache[s]; exists {
			return isPalindrome
		}
		i := 0
		j := len(s) - 1
		for {
			if i >= j {
				break
			}
			if s[i] != s[j] {
				cache[s] = false
				return false
			}
			i += 1
			j -= 1
		}
		cache[s] = true
		return true
	}
}

type Backtrack = func(start int)
type WhenFound = func(partition []string)

func backtrackFunc(s string, isPalindrome IsPalindrome, whenFound WhenFound) Backtrack {
	current := []string{}
	var backtrack Backtrack
	backtrack = func(start int) {
		// 直到所有字元都用盡時，取用當前堆疊。
		if start > len(s)-1 {
			partition := make([]string, len(current))
			copy(partition, current)
			whenFound(partition)
			return
		}
		// 往後找，若找到迴文，將它推入堆疊並繼續往下分解。
		for end := start + 1; end <= len(s); end += 1 {
			str := s[start:end]
			if !isPalindrome(str) {
				continue
			}
			current = append(current, str)
			backtrack(end)
			// 處理完時記得將迴文從堆疊取出
			current = current[:len(current)-1]
		}
	}
	return backtrack
}

/*
Given a string s, partition s such that every
substring of the partition is a palindrome. Return all possible palindrome partitioning of s.
*/
func partition(s string) [][]string {
	partitions := [][]string{}
	if len(s) == 1 {
		partitions = append(partitions, []string{s})
	}
	if len(s) <= 1 {
		return partitions
	}

	backtrack := backtrackFunc(s, isPalindromeFunc(), func(partition []string) {
		partitions = append(partitions, partition)
	})
	backtrack(0)
	return partitions
}
