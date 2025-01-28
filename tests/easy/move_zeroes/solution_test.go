package move_zeroes

//lint:file-ignore ST1001 _

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_moveZeroes_1(t *testing.T) {
	t.SkipNow()
	got := []int{0, 1, 0, 3, 12}
	moveZeroes(got)
	expected := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(got, expected) {
		fmt.Printf("%v\n", got)
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_moveZeroes_2(t *testing.T) {
	t.SkipNow()
	got := []int{0}
	moveZeroes(got)
	expected := []int{0}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
