package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	newNodes := map[*Node]*Node{}
	var newHead *Node = nil
	var lastNode *Node = nil
	ops := []func(){}
	node := head
	for {
		if node == nil {
			break
		}
		newNode := &Node{
			Val: node.Val,
		}
		newNodes[node] = newNode
		if lastNode == nil {
			newHead = newNode
		} else {
			lastNode.Next = newNode
		}
		ops = append(ops, func(node *Node, random *Node) func() {
			return func() {
				// fmt.Printf("%v\n", pointer)
				node.Random = newNodes[random]
			}
		}(newNode, node.Random)) // 必須在執行時複製一份，否則屆時可能已經被 gc
		node = node.Next
		lastNode = newNode
	}
	for _, op := range ops {
		op()
	}
	return newHead
}
