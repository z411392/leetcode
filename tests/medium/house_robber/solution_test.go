package house_robber

import (
	"testing"
)

func Test_rob_1(t *testing.T) {
	// t.SkipNow()
	got := rob([]int{1, 2, 3, 1})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_rob_2(t *testing.T) {
	// t.SkipNow()
	got := rob([]int{2, 7, 9, 3, 1})
	expected := 12
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_rob_3(t *testing.T) {
	// t.SkipNow()
	got := rob([]int{2, 1, 1, 2})
	expected := 4
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
