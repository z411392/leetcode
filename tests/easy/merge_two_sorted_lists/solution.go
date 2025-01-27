package merge_two_sorted_lists

//lint:file-ignore ST1001 _

import (
	"reflect"

	. "github.com/z411392/leetcode/pkg/data_structure/binary_heap"
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

func zip(lists ...*ListNode) <-chan int {
	streamCh := make(chan int)

	go func() {
		defer close(streamCh)

		cases := []reflect.SelectCase{}
		for _, list := range lists {
			if list == nil {
				continue
			}

			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ConvertLinkedListToStream(list)),
			})
		}

		if len(cases) == 0 {
			return
		}

		for len(cases) > 0 {
			chosen, value, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}
			streamCh <- value.Interface().(int) // 將值從 channel 取出並轉型
		}
	}()

	return streamCh
}

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	minHeap := NewBinaryHeap(func(a, b int) bool {
		return a < b
	})
	for value := range zip(list1, list2) {
		minHeap.Insert(value)
	}
	linkedList := NewLinkedListFromStream(minHeap.ToStream())
	return linkedList
}
