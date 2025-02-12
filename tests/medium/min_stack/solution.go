package min_stack

type Node struct {
	val  int
	min  int
	next *Node
}

type MinStack struct {
	head *Node
}

func Constructor() MinStack {
	return MinStack{
		head: nil,
	}
}

func (ms *MinStack) Push(val int) {
	if ms.head == nil {
		ms.head = &Node{
			val:  val,
			min:  val,
			next: nil,
		}
		return
	}

	// 每個節點都保存到該節點為止的最小值
	minVal := val
	if ms.head.min < val {
		minVal = ms.head.min
	}

	ms.head = &Node{
		val:  val,
		min:  minVal,
		next: ms.head,
	}
}

func (ms *MinStack) Pop() {
	if ms.head != nil {
		ms.head = ms.head.next
	}
}

func (ms *MinStack) Top() int {
	if ms.head == nil {
		return 0
	}
	return ms.head.val
}

func (ms *MinStack) GetMin() int {
	if ms.head == nil {
		return 0
	}
	return ms.head.min
}
