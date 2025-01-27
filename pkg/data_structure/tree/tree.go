package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeFromSlice(slice []int, index int, null int) (node *TreeNode) {
	if index >= len(slice) {
		return
	}
	value := slice[index]
	if slice[index] == null {
		return
	}
	node = &TreeNode{Val: value}
	node.Left = NewTreeFromSlice(slice, 2*index+1, null)
	node.Right = NewTreeFromSlice(slice, 2*index+2, null)
	return
}

func ConvertTreeToSlice(root *TreeNode) []int {
	slice := []int{}
	if root == nil {
		return []int{}
	}
	nodes := []*TreeNode{root}
	for {
		next := []*TreeNode{}
		for _, node := range nodes {
			slice = append(slice, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) == 0 {
			break
		}
		nodes = next
	}
	return slice
}
