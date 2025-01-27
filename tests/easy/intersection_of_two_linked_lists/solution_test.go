package intersection_of_two_linked_lists

//lint:file-ignore ST1001 _

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_getIntersectionNode_1(t *testing.T) {
	t.SkipNow()
	var intersection *ListNode = nil
	{
		intersection_3 := &ListNode{
			Val:  5,
			Next: nil,
		}
		intersection_2 := &ListNode{
			Val:  4,
			Next: intersection_3,
		}
		intersection_1 := &ListNode{
			Val:  8,
			Next: intersection_2,
		}
		intersection = intersection_1
	}
	var listA *ListNode = nil
	{
		listANode2 := &ListNode{
			Val:  1,
			Next: intersection,
		}
		listANode1 := &ListNode{
			Val:  4,
			Next: listANode2,
		}
		listA = listANode1
	}
	var listB *ListNode = nil
	{
		listBNode3 := &ListNode{
			Val:  1,
			Next: intersection,
		}
		listBNode2 := &ListNode{
			Val:  6,
			Next: listBNode3,
		}
		listBNode1 := &ListNode{
			Val:  5,
			Next: listBNode2,
		}
		listB = listBNode1
	}
	got := getIntersectionNode(listA, listB)
	expected := intersection
	if got != expected {
		t.FailNow()
	}
}

func Test_getIntersectionNode_2(t *testing.T) {
	t.SkipNow()
	var intersection *ListNode = nil
	{
		intersection_2 := &ListNode{
			Val:  4,
			Next: nil,
		}
		intersection_1 := &ListNode{
			Val:  2,
			Next: intersection_2,
		}
		intersection = intersection_1
	}
	var listA *ListNode = nil
	{
		listANode3 := &ListNode{
			Val:  1,
			Next: intersection,
		}
		listANode2 := &ListNode{
			Val:  9,
			Next: listANode3,
		}
		listANode1 := &ListNode{
			Val:  1,
			Next: listANode2,
		}
		listA = listANode1
	}
	var listB *ListNode = nil
	{
		listBNode1 := &ListNode{
			Val:  3,
			Next: intersection,
		}
		listB = listBNode1
	}
	got := getIntersectionNode(listA, listB)
	expected := intersection
	if got != expected {
		t.FailNow()
	}
}

func Test_getIntersectionNode_3(t *testing.T) {
	t.SkipNow()
	var intersection *ListNode = nil
	var listA *ListNode = nil
	{
		listANode3 := &ListNode{
			Val:  4,
			Next: intersection,
		}
		listANode2 := &ListNode{
			Val:  6,
			Next: listANode3,
		}
		listANode1 := &ListNode{
			Val:  2,
			Next: listANode2,
		}
		listA = listANode1
	}
	var listB *ListNode = nil
	{
		listBNode2 := &ListNode{
			Val:  5,
			Next: intersection,
		}
		listBNode1 := &ListNode{
			Val:  1,
			Next: listBNode2,
		}
		listB = listBNode1
	}
	got := getIntersectionNode(listA, listB)
	expected := intersection
	if got != expected {
		t.FailNow()
	}
}
