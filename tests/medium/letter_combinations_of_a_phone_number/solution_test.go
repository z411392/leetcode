package letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func Test_letterCombinations_1(t *testing.T) {
	t.SkipNow()
	got := letterCombinations("23")
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_letterCombinations_2(t *testing.T) {
	t.SkipNow()
	got := letterCombinations("")
	expected := []string{}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_letterCombinations_3(t *testing.T) {
	t.SkipNow()
	got := letterCombinations("2")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
