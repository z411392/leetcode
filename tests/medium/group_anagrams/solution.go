package group_anagrams

import (
	"strconv"
	"strings"
)

func rleOf(stats [26]int) string {
	stringBuilder := &strings.Builder{}
	for i := 0; i < 26; i += 1 {
		stringBuilder.WriteString(strconv.Itoa(stats[i]))
		stringBuilder.WriteRune(rune('A' + i))
	}
	return stringBuilder.String()
}

func statsOf(str string) [26]int {
	stats := [26]int{}
	for _, character := range str {
		n := character - 'a'
		stats[n] += 1
	}
	return stats
}

var indexMapping = map[string]int{}
var autoIncrement = -1

func indexOf(rle string) int {
	if _, exists := indexMapping[rle]; !exists {
		autoIncrement += 1
		indexMapping[rle] = autoIncrement
	}
	return indexMapping[rle]
}

/*
Given an array of strings strs, group the
anagrams together. You can return the answer in any order.
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		return [][]string{strs}
	}
	groups := [][]string{}
	for _, str := range strs {
		stats := statsOf(str)
		rle := rleOf(stats)
		index := indexOf(rle)
		if index >= len(groups) {
			groups = append(groups, []string{str})
		} else {
			groups[index] = append(groups[index], str)
		}
	}
	return groups
}
