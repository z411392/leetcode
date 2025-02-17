package top_k_frequent_elements

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"
)

func Test_topKFrequent_1(t *testing.T) {
	t.SkipNow()
	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	expected := []int{1, 2}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_topKFrequent_2(t *testing.T) {
	t.SkipNow()
	got := topKFrequent([]int{1}, 1)
	expected := []int{1}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_topKFrequent_3(t *testing.T) {
	// t.SkipNow()
	got := topKFrequent([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 10)
	expected := []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
