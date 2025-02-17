package game_of_life

import (
	"reflect"
	"testing"
)

func Test_gameOfLife_1(t *testing.T) {
	// t.SkipNow()
	got := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(got)
	expected := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_gameOfLife_2(t *testing.T) {
	// t.SkipNow()
	got := [][]int{
		{1, 1},
		{1, 0},
	}
	gameOfLife(got)
	expected := [][]int{
		{1, 1},
		{1, 1},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
