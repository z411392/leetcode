package word_search

import (
	"testing"
)

func Test_exist_1(t *testing.T) {
	// t.SkipNow()
	got := exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED")
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_exist_2(t *testing.T) {
	// t.SkipNow()
	got := exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE")
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
