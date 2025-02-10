package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for {
		nodes := queue
		queue = []*Node{}
		for i := 0; i < len(nodes); i += 1 {
			node := nodes[i]
			if i >= 1 {
				lastNode := nodes[i-1]
				lastNode.Next = node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if len(queue) == 0 {
			break
		}
	}
	return root
}
