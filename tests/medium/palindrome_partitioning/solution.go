package palindrome_partitioning

func getAllPartitions(s string) [][]string {
	partitions := [][]string{}
	if len(s) == 0 {
		return partitions
	}
	if len(s) > 1 {
		for i := 1; i <= len(s); i += 1 {
			one := s[0:i]
			subPartitions := getAllPartitions(s[i:])
			for _, subPartition := range subPartitions {
				partition := append([]string{one}, subPartition...)
				partitions = append(partitions, partition)
			}
		}
	}
	partitions = append(partitions, []string{s})
	return partitions
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return isPalindrome(s[1 : len(s)-1])
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
	for _, partition := range getAllPartitions(s) {
		found := true
		for _, str := range partition {
			// fmt.Printf("%s\n", str)
			if !isPalindrome(str) {
				found = false
				break
			}
		}
		if !found {
			continue
		}
		partitions = append(partitions, partition)
	}
	return partitions
}
