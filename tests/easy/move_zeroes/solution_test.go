package move_zeroes

//lint:file-ignore ST1001 _

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_moveZeroes_1(t *testing.T) {
	t.SkipNow()
	slice := []int{0, 1, 0, 3, 12}
	moveZeroes(slice)
	if !reflect.DeepEqual(slice, []int{1, 3, 12, 0, 0}) {
		fmt.Printf("%v\n", slice)
		t.FailNow()
	}
}

func Test_moveZeroes_2(t *testing.T) {
	t.SkipNow()
	slice := []int{0}
	moveZeroes(slice)
	if !reflect.DeepEqual(slice, []int{0}) {
		t.FailNow()
	}
}
