package unique_paths

import (
	"reflect"
	"testing"
)

func Test_uniquePaths_1(t *testing.T) {
	// t.SkipNow()
	got := uniquePaths(3, 7)
	expected := 28
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_uniquePaths_2(t *testing.T) {
	// t.SkipNow()
	got := uniquePaths(3, 2)
	expected := 3
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
