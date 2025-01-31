package longest_palindromic_substring

import (
	"testing"
)

func Test_longestPalindrome_1(t *testing.T) {
	t.SkipNow()
	got := longestPalindrome("babad")
	expected := "bab"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_longestPalindrome_2(t *testing.T) {
	t.SkipNow()
	got := longestPalindrome("cbbd")
	expected := "bb"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
