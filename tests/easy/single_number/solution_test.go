package single_number

import (
	"testing"
)

func Test_singleNumber_1(t *testing.T) {
	t.SkipNow()
	got := singleNumber([]int{2, 2, 1})
	expected := 1
	if got != expected {
		t.FailNow()
	}
}

func Test_singleNumber_2(t *testing.T) {
	t.SkipNow()
	got := singleNumber([]int{4, 1, 2, 1, 2})
	expected := 4
	if got != expected {
		t.FailNow()
	}
}

func Test_singleNumber_3(t *testing.T) {
	t.SkipNow()
	got := singleNumber([]int{1})
	expected := 1
	if got != expected {
		t.FailNow()
	}
}
