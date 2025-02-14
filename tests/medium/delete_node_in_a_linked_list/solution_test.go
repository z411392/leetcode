package delete_node_in_a_linked_list

//lint:file-ignore ST1001 _

import (
	"reflect"
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_deleteNode_1(t *testing.T) {
	// t.SkipNow()
	head := NewLinkedListFromSlice([]int{4, 5, 1, 9})
	node := head
	value := 5
	for {
		if node.Val == value {
			break
		}
		node = node.Next
	}
	deleteNode(node)
	got := ConvertLinkedListToSlice(head)
	expected := []int{4, 1, 9}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_deleteNode_2(t *testing.T) {
	// t.SkipNow()
	head := NewLinkedListFromSlice([]int{4, 5, 1, 9})
	node := head
	value := 1
	for {
		if node.Val == value {
			break
		}
		node = node.Next
	}
	deleteNode(node)
	got := ConvertLinkedListToSlice(head)
	expected := []int{4, 5, 9}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
