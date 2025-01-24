package linked_list

type LinkedList[T any] struct {
	Val  T
	Next *LinkedList[T]
}
