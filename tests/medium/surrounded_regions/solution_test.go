package surrounded_regions

import (
	"reflect"
	"testing"
)

func Test_solve_1(t *testing.T) {
	// t.SkipNow()
	got := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(got)
	expected := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_solve_2(t *testing.T) {
	// t.SkipNow()
	got := [][]byte{
		{'X'},
	}
	solve(got)
	expected := [][]byte{
		{'X'},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
