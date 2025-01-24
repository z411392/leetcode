package merge_two_sorted_lists

//lint:file-ignore ST1001 Dot imports are used for convenience in this test file

import (
	"reflect"

	. "github.com/z411392/leetcode/pkg/data_structure/binary_heap"
	. "github.com/z411392/leetcode/pkg/data_structure/linked_list"
)

type ListNode = LinkedList[int]

func newLinkedListFromStream[T any](streamCh <-chan T) *LinkedList[T] {
	var head, current *LinkedList[T]

	for value := range streamCh {
		newNode := &LinkedList[T]{Val: value}

		if head == nil {
			head = newNode
		} else {
			current.Next = newNode
		}
		current = newNode
	}

	return head
}

func convertLinkedListToStream[T any](list *LinkedList[T]) <-chan T {
	streamCh := make(chan T)
	if list == nil {
		close(streamCh)
		return streamCh
	}

	go func() {
		defer close(streamCh)

		current := list
		for current != nil {
			streamCh <- current.Val
			current = current.Next
		}
	}()

	return streamCh
}

func merge[T any](lists ...*LinkedList[T]) <-chan T {
	streamCh := make(chan T)

	go func() {
		defer close(streamCh)

		cases := []reflect.SelectCase{}
		for _, list := range lists {
			if list == nil {
				continue
			}

			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(convertLinkedListToStream(list)),
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
			streamCh <- value.Interface().(T)
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
	for value := range merge(list1, list2) {
		minHeap.Insert(value)
	}
	linkedList := newLinkedListFromStream(minHeap.ToStream())
	return linkedList
}
