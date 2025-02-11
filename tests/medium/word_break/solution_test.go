package word_break

import (
	"testing"
)

func Test_wordBreak_1(t *testing.T) {
	t.SkipNow()
	got := wordBreak("leetcode", []string{"leet", "code"})
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_wordBreak_2(t *testing.T) {
	t.SkipNow()
	got := wordBreak("applepenapple", []string{"apple", "pen"})
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_wordBreak_3(t *testing.T) {
	t.SkipNow()
	got := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_wordBreak_4(t *testing.T) {
	t.SkipNow()
	got := wordBreak("aaaaaaa", []string{"aaaa", "aaa"})
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_wordBreak_5(t *testing.T) {
	// t.SkipNow()
	got := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"})
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
