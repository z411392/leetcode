package jump_game

import (
	"testing"
)

func Test_canJump_1(t *testing.T) {
	t.SkipNow()
	got := canJump([]int{2, 3, 1, 1, 4})
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_canJump_2(t *testing.T) {
	t.SkipNow()
	got := canJump([]int{3, 2, 1, 0, 4})
	expected := false
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
