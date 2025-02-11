package palindrome_partitioning

import (
	"reflect"
	"testing"
)

func Test_partition_1(t *testing.T) {
	t.SkipNow()
	got := partition("aab")
	expected := [][]string{
		{"a", "a", "b"},
		{"aa", "b"},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_partition_2(t *testing.T) {
	t.SkipNow()
	got := partition("a")
	expected := [][]string{
		{"a"},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_partition_3(t *testing.T) {
	// t.SkipNow()
	got := partition("bb")
	expected := [][]string{
		{"b", "b"},
		{"bb"},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
