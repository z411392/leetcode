package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func iterateOnLevel(node *Node, fun func(node *Node)) {
	for {
		if node == nil {
			break
		}
		fun(node)
		node = node.Next
	}
}

func connect(root *Node) *Node {
	node := root
	for {
		// 因為總是在這一層處理下一層的 next 關係。所以如果沒有下一層，就可以提早結束了。
		if node == nil || node.Left == nil {
			break
		}
		iterateOnLevel(node, func(node *Node) {
			// 除了將左節點的 Next 設為右節點外，還要將右節點的 Next 設為 Next 的左節點。
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		})
		node = node.Left
	}
	return root
}
