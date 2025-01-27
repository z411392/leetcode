package palindrome_linked_list

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_isPalindrome_1(t *testing.T) {
	t.SkipNow()
	got := isPalindrome(NewLinkedListFromSlice([]int{1, 2, 2, 1}))
	expected := true
	if got != expected {
		t.FailNow()
	}
}

func Test_isPalindrome_2(t *testing.T) {
	t.SkipNow()
	got := isPalindrome(NewLinkedListFromSlice([]int{1, 2}))
	expected := false
	if got != expected {
		t.FailNow()
	}
}
