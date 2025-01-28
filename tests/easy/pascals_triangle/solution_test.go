package pascals_triangle

import (
	"reflect"
	"testing"
)

func Test_generate_1(t *testing.T) {
	t.SkipNow()
	got := generate(5)
	expected := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_generate_2(t *testing.T) {
	t.SkipNow()
	got := generate(1)
	expected := [][]int{{1}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
