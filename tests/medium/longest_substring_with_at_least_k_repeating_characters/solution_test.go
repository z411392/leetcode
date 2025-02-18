package longest_substring_with_at_least_k_repeating_characters

import (
	"testing"
)

func Test_longestSubstring_1(t *testing.T) {
	// t.SkipNow()
	got := longestSubstring("aaabb", 3)
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_longestSubstring_2(t *testing.T) {
	// t.SkipNow()
	got := longestSubstring("ababbc", 2)
	expected := 5
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
