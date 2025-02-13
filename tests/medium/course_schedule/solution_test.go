package course_schedule

import (
	"testing"
)

func Test_canFinish_1(t *testing.T) {
	// t.SkipNow()
	got := canFinish(2, [][]int{{1, 0}})
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_canFinish_2(t *testing.T) {
	// t.SkipNow()
	got := canFinish(2, [][]int{{1, 0}, {0, 1}})
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
