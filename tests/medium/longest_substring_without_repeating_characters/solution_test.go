package longest_substring_without_repeating_characters

//lint:file-ignore ST1001 _

import (
	"testing"
)

func Test_lengthOfLongestSubstring_1(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLongestSubstring("abcabcbb")
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_lengthOfLongestSubstring_2(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLongestSubstring("bbbbb")
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_lengthOfLongestSubstring_3(t *testing.T) {
	// t.SkipNow()
	got := lengthOfLongestSubstring("pwwkew")
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
