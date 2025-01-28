package best_time_to_buy_and_sell_stock

import (
	"testing"

	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func Test_hasCycle_1(t *testing.T) {
	t.SkipNow()
	var list *ListNode = nil
	{
		node4 := &ListNode{
			Val:  -4,
			Next: nil,
		}
		node3 := &ListNode{
			Val:  0,
			Next: node4,
		}
		node2 := &ListNode{
			Val:  2,
			Next: node3,
		}
		node1 := &ListNode{
			Val:  3,
			Next: node2,
		}
		node4.Next = node2
		list = node1
	}
	got := hasCycle(list)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}

func Test_hasCycle_2(t *testing.T) {
	t.SkipNow()
	var list *ListNode = nil
	{
		node2 := &ListNode{
			Val:  2,
			Next: nil,
		}
		node1 := &ListNode{
			Val:  1,
			Next: node2,
		}
		node2.Next = node1
		list = node1
	}
	got := hasCycle(list)
	expected := true
	if got != expected {
		t.Fatalf("expected %v, got %v\n", expected, got)
	}
}
