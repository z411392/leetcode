package count_and_say

import (
	"testing"
)

func Test_countAndSay_1(t *testing.T) {
	// t.SkipNow()
	got := countAndSay(4)
	expected := "1211"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_countAndSay_2(t *testing.T) {
	// t.SkipNow()
	got := countAndSay(1)
	expected := "1"
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
