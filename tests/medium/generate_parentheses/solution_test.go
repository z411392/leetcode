package generate_parentheses

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis_1(t *testing.T) {
	t.SkipNow()
	got := generateParenthesis(3)
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_generateParenthesis_2(t *testing.T) {
	t.SkipNow()
	got := generateParenthesis(1)
	expected := []string{"()"}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
