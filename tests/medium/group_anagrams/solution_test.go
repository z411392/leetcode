package group_anagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams_1(t *testing.T) {
	// t.SkipNow()
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_groupAnagrams_2(t *testing.T) {
	// t.SkipNow()
	got := groupAnagrams([]string{""})
	expected := [][]string{{""}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_groupAnagrams_3(t *testing.T) {
	// t.SkipNow()
	got := groupAnagrams([]string{"a"})
	expected := [][]string{{"a"}}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
