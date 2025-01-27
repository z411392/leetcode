package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedListFromSlice(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	head := &ListNode{Val: slice[0]}
	current := head

	for _, value := range slice[1:] {
		current.Next = &ListNode{Val: value}
		current = current.Next
	}

	return head
}

func ConvertLinkedListToSlice(list *ListNode) []int {
	if list == nil {
		return []int{}
	}

	slice := []int{}
	current := list
	for current != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}

	return slice
}

func NewLinkedListFromStream(streamCh <-chan int) *ListNode {
	var head, current *ListNode

	for value := range streamCh {
		newNode := &ListNode{Val: value}

		if head == nil {
			head = newNode
		} else {
			current.Next = newNode
		}
		current = newNode
	}

	return head
}

func ConvertLinkedListToStream(list *ListNode) <-chan int {
	streamCh := make(chan int)
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
