package longest_common_prefix

import (
	"testing"
)

func Test_longestCommonPrefix_1(t *testing.T) {
	t.SkipNow()
	got := longestCommonPrefix([]string{"flower", "flow", "flight"})
	expected := "fl"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_longestCommonPrefix_2(t *testing.T) {
	t.SkipNow()
	got := longestCommonPrefix([]string{"dog", "racecar", "car"})
	expected := ""
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
