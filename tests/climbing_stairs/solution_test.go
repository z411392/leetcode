package climbing_stairs

import (
	"testing"
)

func Test_climbStairs_1(t *testing.T) {
	t.SkipNow()
	got := climbStairs(2)
	expected := 2
	if got != expected {
		t.FailNow()
	}
}

func Test_climbStairs_2(t *testing.T) {
	t.SkipNow()
	got := climbStairs(3)
	expected := 3
	if got != expected {
		t.FailNow()
	}
}

func Test_climbStairs_3(t *testing.T) {
	t.SkipNow()
	got := climbStairs(35)
	expected := 14930352
	if got != expected {
		t.FailNow()
	}
}
