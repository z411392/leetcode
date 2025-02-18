package longest_substring_with_at_least_k_repeating_characters

func longestSubstring(s string, k int) int {
	counters := map[rune]int{}
	for _, c := range s {
		counters[c] += 1
	}
	removing := map[rune]int{}
	for c, counter := range counters {
		if counter >= k {
			continue
		}
		removing[c] = counter
	}
	left := 0
	right := len(s) - 1
	for {
		if len(removing) == 0 {
			break
		}
		// fmt.Printf("%s\n", s[left:right+1])
		if left >= right {
			break
		}
		// time.Sleep(time.Duration(500 * time.Millisecond))
		found := false
		fromLeft := rune(s[left])
		_, exists := removing[fromLeft]
		if exists {
			found = true
			left += 1
			counters[fromLeft] -= 1
			if counters[fromLeft] == 0 {
				delete(removing, fromLeft)
			}
		}
		fromRight := rune(s[right])
		_, exists = removing[fromRight]
		if exists {
			found = true
			right -= 1
			counters[fromRight] -= 1
			if counters[fromRight] == 0 {
				delete(removing, fromRight)
			}
		}
		if !found {
			if counters[fromLeft] > counters[fromRight] {
				counters[fromLeft] -= 1
				if counters[fromLeft] < k {
					removing[fromLeft] = counters[fromLeft]
				}
				left += 1
			} else {
				counters[fromRight] -= 1
				if counters[fromRight] < k {
					removing[fromRight] = counters[fromRight]
				}
				right -= 1
			}
		}
	}
	if left == right {
		if k == 1 {
			return 1
		}
		return 0
	}
	return len(s[left : right+1])
}
