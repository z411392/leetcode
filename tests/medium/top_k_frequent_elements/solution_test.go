package top_k_frequent_elements

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"
)

func Test_topKFrequent_1(t *testing.T) {
	// t.SkipNow()
	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	expected := []int{1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_topKFrequent_2(t *testing.T) {
	// t.SkipNow()
	got := topKFrequent([]int{1}, 1)
	expected := []int{1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
