package valid_parentheses

import (
	"testing"
)

func Test_isValid_1(t *testing.T) {
	t.SkipNow()
	got := isValid("()")
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isValid_2(t *testing.T) {
	t.SkipNow()
	got := isValid("()[]{}")
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isValid_3(t *testing.T) {
	t.SkipNow()
	got := isValid("(]")
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_isValid_4(t *testing.T) {
	t.SkipNow()
	got := isValid("([])")
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
