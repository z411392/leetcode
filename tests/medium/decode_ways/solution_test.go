package decode_ways

import (
	"testing"
)

func Test_numDecodings_1(t *testing.T) {
	t.SkipNow()
	got := numDecodings("12")
	expected := 2
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_numDecodings_2(t *testing.T) {
	t.SkipNow()
	got := numDecodings("226")
	expected := 3
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_numDecodings_3(t *testing.T) {
	t.SkipNow()
	got := numDecodings("06")
	expected := 0
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_numDecodings_4(t *testing.T) {
	t.SkipNow()
	got := numDecodings("2101")
	expected := 1
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
