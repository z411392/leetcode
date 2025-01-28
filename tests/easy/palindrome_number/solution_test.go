package palindrome_number

import (
	"testing"
)

func Test_isPalindrome_1(t *testing.T) {
	t.SkipNow()
	got := isPalindrome(121)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isPalindrome_2(t *testing.T) {
	t.SkipNow()
	got := isPalindrome(-121)
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isPalindrome_3(t *testing.T) {
	t.SkipNow()
	got := isPalindrome(10)
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
